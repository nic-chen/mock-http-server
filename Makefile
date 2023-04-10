VERSION ?= v1

IMAGE_REPOSITORY ?= "johz/mock-server"
IMAGE_TAG = $(VERSION)

GOOS ?= linux
GOARCH ?= amd64

### function for get build os and arch args
### param: $(1) OS
### param: $(2) ARCH
go_build_args = GOOS=$(1) GOARCH=$(2)


.PHONY: help
help:  ## Display this help message.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build:  ## Build the binary file.
	$(call go_build_args,$(GOOS),$(GOARCH)) CGO_ENABLED=0 go build -o ./mock-server -tags $(VERSION) .

.PHONY: build-docker-image
build-docker-image: build ## Build the docker image.
	@docker build -t $(IMAGE_REPOSITORY):$(IMAGE_TAG) .
	@rm ./mock-server
