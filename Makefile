ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack.mk

tool:
	go build -o ai tools/run/main.go  && mv ai ${GOPATH}/bin
	make db

pb:
	json=on && protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/gate.proto
	json=on && protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/db.proto
	json=on && protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/struct.proto
	json=on && protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/enum.proto
	json=on && protoc --proto_path=manifest/protobuf/proto --gogofast_out=paths=source_relative:manifest/protobuf/pb manifest/protobuf/proto/gpt.proto
db:pb
	cd manifest/protobuf && ai db

go:db
	#编译服务
	protoc --proto_path=manifest/protobuf/srv --proto_path=manifest/protobuf/proto --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:manifest/protobuf/pb manifest/protobuf/srv/aisrv.proto

aisrv:
