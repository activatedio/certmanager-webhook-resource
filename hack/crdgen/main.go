package main

import (
	"github.com/activatedio/certmanager-webhook-resource/pkg/crd"
	"github.com/activatedio/wrangler/crdgen"
)

func main() {
	crdgen.Run(crd.List())
}
