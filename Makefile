#
# Makefile for building, running, and testing
#

# Import dotenv
ifneq (,$(wildcard ../.env))
	include ../.env
	export
endif

# Application versions
BASE_VERSION = $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH = $(shell git rev-parse --short HEAD)


# Gets the directory containing the Makefile
ROOT_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Base container registry
SRO_BASE_REGISTRY ?= 779965382548.dkr.ecr.us-east-1.amazonaws.com
SRO_REGISTRY ?= $(SRO_BASE_REGISTRY)/sro

# The registry for this service
REGISTRY = $(SRO_REGISTRY)/accounts
time=$(shell date +%s)

PROTO_DIR=$(ROOT_DIR)/api
PROTO_THIRD_PARTY_DIR=$(ROOT_DIR)/third_party

PROTO_FILES = $(shell find $(PROTO_DIR) -name '*.proto')

test:
	ginkgo $(ROOT_DIR)/... -covermode atomic

report: test
	go tool cover -func=$(ROOT_DIR)/coverage.out
	go tool cover -html=$(ROOT_DIR)/coverage.out

build: build-character build-chat build-gamebackend
build-%:
	go build -ldflags="-X 'github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/config/default.Version=$(BASE_VERSION)'" -o $(ROOT_DIR)/bin/$* $(ROOT_DIR)/cmd/$*  

run: run-character run-chat run-gamebackend
run-%:
	go run $(ROOT_DIR)/cmd/$*

deploy: aws-docker-login push

buildi: buildi-character buildi-chat buildi-gamebackend
buildi-%:
	docker build --build-arg APP_VERSION=$(BASE_VERSION) -t sro-$* -f build/$*.Dockerfile .

aws-docker-login:
	aws ecr get-login-password | docker login --username AWS --password-stdin $(SRO_BASE_REGISTRY)

pushf: pushf-character pushf-chat pushf-gamebackend
pushf-%:
	docker tag sro-$* $(SRO_REGISTRY)/$*:latest
	docker tag sro-$* $(SRO_REGISTRY)/$*:$(BASE_VERSION)
	docker tag sro-$* $(SRO_REGISTRY)/$*:$(BASE_VERSION)-$(COMMIT_HASH)
	docker push $(SRO_REGISTRY)/$*:latest
	docker push $(SRO_REGISTRY)/$*:$(BASE_VERSION)
	docker push $(SRO_REGISTRY)/$*:$(BASE_VERSION)-$(COMMIT_HASH)

push: push-character push-chat push-gamebackend
push-%: buildi-%
	docker tag sro-$* $(SRO_REGISTRY)/$*:latest
	docker tag sro-$* $(SRO_REGISTRY)/$*:$(BASE_VERSION)
	docker tag sro-$* $(SRO_REGISTRY)/$*:$(BASE_VERSION)-$(COMMIT_HASH)
	docker push $(SRO_REGISTRY)/$*:latest
	docker push $(SRO_REGISTRY)/$*:$(BASE_VERSION)
	docker push $(SRO_REGISTRY)/$*:$(BASE_VERSION)-$(COMMIT_HASH)

.PHONY: clean-protos protos $(PROTO_FILES)

clean-protos:
	rm -rf "$(ROOT_DIR)/pkg/pb"

protos: clean-protos $(PROTO_FILES)

$(PROTO_FILES):
	protoc "$@" \
		-I "$(PROTO_DIR)" \
		-I "$(PROTO_THIRD_PARTY_DIR)" \
		--go_out="$(ROOT_DIR)" \
		--go-grpc_out="$(ROOT_DIR)" \
		--grpc-gateway_out="$(ROOT_DIR)" \
		--grpc-gateway_opt logtostderr=true 

download:
	go mod download

install-tools:
	  cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
