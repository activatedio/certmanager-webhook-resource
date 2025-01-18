package main

import (
	"github.com/activatedio/certmanager-webhook-resource/pkg/crd"
	"os"

	_ "github.com/activatedio/certmanager-webhook-resource/pkg/apis/resource.cert-manager.activated.io/v1"
	wcrd "github.com/rancher/wrangler/v3/pkg/crd"
)

//go:generate go run .
func main() {
	wcrd.Print(os.Stdout, crd.List())
}
