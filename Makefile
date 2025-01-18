GIT_BRANCH?=$(shell git branch --show-current)
GIT_COMMIT?=$(shell git rev-parse HEAD)
GIT_COMMIT_SHORT?=$(shell git rev-parse --short HEAD)
GIT_TAG?=v0.0.0
ifneq ($(GIT_BRANCH), main)
GIT_TAG?=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "v0.0.0" )
endif
TAG?=${GIT_TAG}-${GIT_COMMIT_SHORT}
CHART_VERSION?=900 # Only used in e2e to avoid downgrades from rancher

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR := $(abspath $(ROOT_DIR)/bin)


ifdef GOROOT
GOROOT := ${GOROOT}
else
GOROOT := ${HOME}/go
endif

default: operator

.PHONY: generate
generate:
	go generate gen.go
	go install golang.org/x/tools/cmd/goimports
	${GOROOT}/bin/goimports -w ./pkg
	go generate ./hack/crdgen.go > charts/certmanager-webhook-resource-crds/templates/crds.yaml

.PHONY: test
test:
	go test ./pkg/...

.PHONY: clean
clean:
	rm -rf build bin dist

.PHONY: operator-chart
webhook-chart:
	mkdir -p $(BIN_DIR)
	cp -rf $(ROOT_DIR)/charts/certmanager-webhook-resource $(BIN_DIR)/chart
	sed -i -e 's/tag:.*/tag: '${TAG}'/' $(BIN_DIR)/chart/values.yaml
	helm package --version ${CHART_VERSION} --app-version ${GIT_TAG} -d $(BIN_DIR)/ $(BIN_DIR)/chart
	rm -Rf $(BIN_DIR)/chart
	
.PHONY: crd-chart
crd-chart:
	mkdir -p $(BIN_DIR)
	helm package --version ${CHART_VERSION} --app-version ${GIT_TAG} -d $(BIN_DIR)/ $(ROOT_DIR)/charts/certmanager-webhook-resource-crds
	rm -Rf $(BIN_DIR)/chart

.PHONY: charts
charts:
	$(MAKE) operator-chart
	$(MAKE) crd-chart

.PHONY: nix-shell
nix-shell:
	nix-shell default.nix --command $${SHELL}

