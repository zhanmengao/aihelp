ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack.mk

tool:
	go build -o ai tools/run/main.go  && mv ai ${GOPATH}/bin
	make db

pb:
	protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/gate.proto
	protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/db.proto
	protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/struct.proto
db:pb
	cd manifest/protobuf && ai db

go:db


aisrv:
