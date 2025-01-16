package main

import (
	"context"
	bluebirdv1 "github.com/activatedio/acre-bluebird-operator/pkg/apis/bluebird.acresecurity.com/v1"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"strings"
	"sync"

	bluebirdclientv1 "github.com/activatedio/acre-bluebird-operator/pkg/generated/clientset/versioned"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

const GroupName = "certmanager.bluebird.acresecurity.com"

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	cmd.RunWebhookServer(GroupName,
		&solver{},
	)
}

type solver struct {
	lock   sync.Mutex
	client *bluebirdclientv1.Clientset
}

type customDNSProviderConfig struct {
	// Change the two fields below according to the format of the configuration
	// to be decoded.
	// These fields will be set by users in the
	// `issuer.spec.acme.dns01.providers.webhook.config` field.

	//Email           string `json:"email"`
	//APIKeySecretRef v1alpha1.SecretKeySelector `json:"apiKeySecretRef"`
}

func (s *solver) Name() string {
	return "resource"
}

func (s *solver) makeName(ch *v1alpha1.ChallengeRequest) string {
	return strings.ReplaceAll(strings.ReplaceAll(ch.DNSName, "*", "wild"), ".", "-")
}

func (s *solver) Present(ch *v1alpha1.ChallengeRequest) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	logrus.Infof("processng dns request %v", ch)

	ctx := context.Background()

	name := s.makeName(ch)

	api := s.client.BluebirdV1().BluebirdDNSChallenges(ch.ResourceNamespace)

	c, err := api.Get(ctx, name, metav1.GetOptions{})

	spec := bluebirdv1.BluebirdDNSChallengeSpec{
		DNSName:      ch.DNSName,
		Key:          ch.Key,
		ResolvedFQDN: ch.ResolvedFQDN,
		ResolvedZone: ch.ResolvedZone,
	}

	if err != nil {
		if apierrors.IsNotFound(err) {

			c = &bluebirdv1.BluebirdDNSChallenge{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: ch.ResourceNamespace,
				},
				Spec: spec,
			}

			logrus.Infof("creating dns challenge %s", name)
			_, err = api.Create(ctx, c, metav1.CreateOptions{})
			return err

		} else {
			return err
		}
	} else {
		c.Spec = spec
		logrus.Infof("updating dns challenge %s", name)
		_, err = api.Update(ctx, c, metav1.UpdateOptions{})
		return err
	}

}

func (s *solver) CleanUp(ch *v1alpha1.ChallengeRequest) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	cts := context.Background()

	name := s.makeName(ch)

	api := s.client.BluebirdV1().BluebirdDNSChallenges(ch.ResourceNamespace)

	logrus.Infof("removing dns challenge %s", name)
	return api.Delete(cts, name, metav1.DeleteOptions{})
}

func (s *solver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {

	cs, err := bluebirdclientv1.NewForConfig(kubeClientConfig)

	if err != nil {
		return err
	}

	s.client = cs

	return nil
}
