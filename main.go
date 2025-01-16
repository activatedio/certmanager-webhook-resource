package main

import (
	"encoding/json"
	"fmt"
	extapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/rest"

	"github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"
)

const GroupName = "certmanager.bluebird.acresecurity.com"

func main() {
	fmt.Println("Running in MAIN!!!")
	cmd.RunWebhookServer(GroupName,
		&solver{},
	)
}

type solver struct {
	// If a Kubernetes 'clientset' is needed, you must:
	// 1. uncomment the additional `client` field in this structure below
	// 2. uncomment the "k8s.io/client-go/kubernetes" import at the top of the file
	// 3. uncomment the relevant code in the Initialize method below
	// 4. ensure your webhook's service account has the required RBAC role
	//    assigned to it for interacting with the Kubernetes APIs you need.
	//client kubernetes.Clientset
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

func (s *solver) Present(ch *v1alpha1.ChallengeRequest) error {
	fmt.Println("In present *******")
	cfg, err := loadConfig(ch.Config)
	if err != nil {
		return err
	}

	fmt.Printf("Decoded configuration %v\n", cfg)

	fmt.Printf("Request %v\n", ch)

	return nil
}

func (s *solver) CleanUp(ch *v1alpha1.ChallengeRequest) error {
	// TODO: add code that deletes a record from the DNS provider's console
	return nil
}

func (s *solver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	fmt.Println("Initialize !!! *******")
	///// UNCOMMENT THE BELOW CODE TO MAKE A KUBERNETES CLIENTSET AVAILABLE TO
	///// YOUR CUSTOM DNS PROVIDER

	//cl, err := kubernetes.NewForConfig(kubeClientConfig)
	//if err != nil {
	//	return err
	//}
	//
	//s.client = cl

	///// END OF CODE TO MAKE KUBERNETES CLIENTSET AVAILABLE
	return nil
}

func loadConfig(cfgJSON *extapi.JSON) (customDNSProviderConfig, error) {
	cfg := customDNSProviderConfig{}
	// handle the 'base case' where no configuration has been provided
	if cfgJSON == nil {
		return cfg, nil
	}
	if err := json.Unmarshal(cfgJSON.Raw, &cfg); err != nil {
		return cfg, fmt.Errorf("error decoding solver config: %v", err)
	}

	return cfg, nil
}
