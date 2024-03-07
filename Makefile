ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack.mk

pb:
	protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/gate.proto