# Global
ORG ?= marceloboeira
PROVIDER ?= flagr
VERSION=1.0.0
BINARY ?= terraform-provider-${PROVIDER}_${VERSION}

# Local Binary (Plugin cache)
LOCAL_BIN ?= bin/${BINARY}
ARCHITECTURE ?= darwin_amd64
PLUGINS_PATH ?= ~/.terraform.d/plugins
PROVIDER_PATH ?= ${PLUGINS_PATH}/registry.terraform.io/${ORG}/${PROVIDER}/${VERSION}/${ARCHITECTURE}

# Release
RELEASE_BIN ?= bin/${VERSION}

## Code
GO_FILES ?= $$(find . -name '*.go')
TF_EXAMPLES_PATH ?= examples/

## Docker
DOCKER ?= docker/
DOCKER_COMPOSE_FILE ?= docker-compose.yaml


## Tests
TEST ?= $$(go list ./... | grep -v 'vendor')

default: install

.PHONY: help
help: ## Lists the available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: format
format: ## Formats go and terraform code
	@gofmt -w ${GO_FILES}
	@terraform fmt -recursive ${TF_EXAMPLES_PATH}

.PHONY: format build
build: ## Builds the local architecture binary to the root folder
	@go build -o ${LOCAL_BIN}

.PHONY: release
release: ## Builds release-binaries for all architectures
	mkdir -p ${RELEASE_BIN}
	GOOS=darwin GOARCH=amd64  go build -o ${RELEASE_BIN}/${BINARY}_darwin_amd64
	GOOS=freebsd GOARCH=386   go build -o ${RELEASE_BIN}/${BINARY}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ${RELEASE_BIN}/${BINARY}_freebsd_amd64
	GOOS=freebsd GOARCH=arm   go build -o ${RELEASE_BIN}/${BINARY}_freebsd_arm
	GOOS=linux GOARCH=386     go build -o ${RELEASE_BIN}/${BINARY}_linux_386
	GOOS=linux GOARCH=amd64   go build -o ${RELEASE_BIN}/${BINARY}_linux_amd64
	GOOS=linux GOARCH=arm     go build -o ${RELEASE_BIN}/${BINARY}_linux_arm
	GOOS=openbsd GOARCH=386   go build -o ${RELEASE_BIN}/${BINARY}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ${RELEASE_BIN}/${BINARY}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ${RELEASE_BIN}/${BINARY}_solaris_amd64
	GOOS=windows GOARCH=386   go build -o ${RELEASE_BIN}/${BINARY}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ${RELEASE_BIN}/${BINARY}_windows_amd64

.PHONY: install
install: build ## Builds the local architecture binary and install it on the local terraform cache
	@mkdir -p ${PROVIDER_PATH}
	@cp ${LOCAL_BIN} ${PROVIDER_PATH}/${BINARY}

.PHONY: test
test: ## Runs tests
	@go test -i ${TEST} || exit 1
	@echo ${TEST} | xargs -t -n4 go test ${TESTARGS} -timeout=30s -parallel=4

.PHONY: testacc
testacc: ## Runs acceptance tests
	@TF_ACC=1 go test ${TEST} -v ${TESTARGS} -timeout 120m

.PHONE: compose
compose: ## Starts test dependencies with docker-compose
	@docker compose -f ${DOCKER}${DOCKER_COMPOSE_FILE} up
