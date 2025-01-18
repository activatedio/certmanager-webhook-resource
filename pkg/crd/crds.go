package crd

import (
	v1 "github.com/activatedio/certmanager-webhook-resource/pkg/apis/resource.cert-manager.activated.io/v1"
	"github.com/rancher/wrangler/v3/pkg/crd"
)

func List() []crd.CRD {
	dnsChallenge := crd.NamespacedType("ACMEDNSChallenge.resource.cert-manager.activated.io/v1").
		WithSchemaFromStruct(v1.ACMEDNSChallenge{}).
		WithColumn("DNSName", ".spec.dnsName").
		WithColumn("Key", ".spec.key").
		WithColumn("ResolvedFQDN", ".spec.resolvedFQDN").
		WithStatus()

	return []crd.CRD{dnsChallenge}
}
