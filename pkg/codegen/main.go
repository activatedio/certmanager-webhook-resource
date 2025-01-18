package main

import (
	v1 "github.com/activatedio/certmanager-webhook-resource/pkg/apis/resource.cert-manager.activated.io/v1"
	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
)

func main() {

	controllergen.Run(args.Options{
		OutputPackage: "github.com/activatedio/certmanager-webhook-resource/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"resource.cert-manager.activated.io": {
				PackageName: "resource.cert-manager.activated.io",
				Types: []interface{}{
					&v1.ACMEDNSChallenge{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
		},
	})

}
