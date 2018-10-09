BUILDCOMMIT := $(shell git describe --dirty --always)
BUILDDATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VER_FLAGS=-X main.commit=$(BUILDCOMMIT) -X main.date=$(BUILDDATE)

.DEFAULT_GOAL:=help

##@ Dependencies

.PHONY: install-build-deps
install-build-deps: ## Install dependencies (packages and tools)
	./hack/install_deps.sh

##@ Build

.PHONY: build
build: ## Build the operator
	go build -ldflags "$(VER_FLAGS)" ./cmd/serverless-operator

##@ Release

.PHONY: release
release: ## Create a release
	goreleaser --rm-dist

.PHONY: dev-release
dev-release: ## Create a development release
	goreleaser --rm-dist --snapshot --skip-publish

##@ Testing & CI

.PHONY: test
test: ## Run unit tests
	@go test -v -covermode=count -coverprofile=coverage.out ./pkg/... ./cmd/...
	@test -z $(COVERALLS_TOKEN) || $(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=circle-ci

.PHONY: lint
lint: ## Run linting over the codebase
	./bin/golangci-lint run

.PHONY: ci
ci: test lint ## Target for CI system to invoke to run tests and linting

##@ Code Generation

.PHONY: codegen
codegen:
	./hack/update-codegen.sh

##@ Utility

.PHONY: help
help:  ## Display this help. Thanks to https://suva.sh/posts/well-documented-makefiles/
@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
