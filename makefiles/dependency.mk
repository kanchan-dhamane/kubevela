
GOLANGCILINT_VERSION ?= v1.38.0

.PHONY: golangci
golangci:
ifneq ($(shell which golangci-lint),)
	@$(OK) golangci-lint is already installed
GOLANGCILINT=$(shell which golangci-lint)
else ifeq (, $(shell which $(GOBIN)/golangci-lint))
	@{ \
	set -e ;\
	echo 'installing golangci-lint-$(GOLANGCILINT_VERSION)' ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) $(GOLANGCILINT_VERSION) ;\
	echo 'Successfully installed' ;\
	}
GOLANGCILINT=$(GOBIN)/golangci-lint
else
	@$(OK) golangci-lint is already installed
GOLANGCILINT=$(GOBIN)/golangci-lint
endif

.PHONY: staticchecktool
staticchecktool:
ifeq (, $(shell which staticcheck))
	@{ \
	set -e ;\
	echo 'installing honnef.co/go/tools/cmd/staticcheck ' ;\
	GO111MODULE=off go get honnef.co/go/tools/cmd/staticcheck ;\
	}
STATICCHECK=$(GOBIN)/staticcheck
else
STATICCHECK=$(shell which staticcheck)
endif

.PHONY: goimports
goimports:
ifeq (, $(shell which goimports))
	@{ \
	set -e ;\
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports ;\
	}
GOIMPORTS=$(GOBIN)/goimports
else
GOIMPORTS=$(shell which goimports)
endif

.PHONY: installcue
installcue:
ifeq (, $(shell which cue))
	@{ \
	set -e ;\
	GO111MODULE=off go get -u cuelang.org/go/cmd/cue ;\
	}
CUE=$(GOBIN)/cue
else
CUE=$(shell which cue)
endif

KUSTOMIZE_VERSION ?= 3.8.2

.PHONY: kustomize
kustomize:
ifeq (, $(shell kustomize version | grep $(KUSTOMIZE_VERSION)))
	@{ \
	set -eo pipefail ;\
	echo 'installing kustomize-v$(KUSTOMIZE_VERSION) into $(GOBIN)' ;\
	curl -sS https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh | bash -s $(KUSTOMIZE_VERSION) $(GOBIN);\
	echo 'Install succeed' ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif