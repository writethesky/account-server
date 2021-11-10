
BASIC_FETCH_SOURCE ?= remote:git@github.com:writethesky/basic-proto.git
#BASIC_FETCH_SOURCE ?= local:../basic-proto
BASIC_FETCH_SOURCE_TYPE=$(shell echo ${BASIC_FETCH_SOURCE}|awk -F ':' '{print $$1}')
BASIC_FETCH_SOURCE_SITE=$(shell echo ${BASIC_FETCH_SOURCE}|awk -F '${BASIC_FETCH_SOURCE_TYPE}:' '{print $$2}')
FETCH_SOURCE_TYPE_REMOTE="remote"
FETCH_SOURCE_TYPE_LOCAL="local"
PROTO_DIRECTORY="proto"
PROTO_TARGET_DIRECTORY="pb"
BUF_VERSION:=1.0.0-rc6

install-tools:
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
    	chmod +x "$(shell go env GOPATH)/bin/buf"
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.6.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

clean:
	rm -rf ${PROTO_DIRECTORY} ${PROTO_TARGET_DIRECTORY}
fetch-proto: clean
	mkdir ${PROTO_DIRECTORY}
	@echo {\"type\": \"${BASIC_FETCH_SOURCE_TYPE}\", \"site\": \"${BASIC_FETCH_SOURCE_SITE}\"}
ifeq ($(BASIC_FETCH_SOURCE_TYPE),$(shell echo $(FETCH_SOURCE_TYPE_REMOTE)))
	@echo "Ready to pull the remote code..."
	git clone ${BASIC_FETCH_SOURCE_SITE} $(PROTO_DIRECTORY)/basic
else ifeq ($(BASIC_FETCH_SOURCE_TYPE),$(shell echo $(FETCH_SOURCE_TYPE_LOCAL)))
	@echo "Ready to copy local code..."
	cp -r ${BASIC_FETCH_SOURCE_SITE} $(PROTO_DIRECTORY)/basic
endif
generate: fetch-proto
	buf generate ${PROTO_DIRECTORY}

run:
	swag init --output ./docs/
	go run main.go

