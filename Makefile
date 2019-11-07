regenerate:
	#(protoc-gen-combo --version="3.0.0"  --proto_path=$(GOPATH)/src/github.com/gogo/protobuf:$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf:.    --gogo_out=. proto.proto)
	(protoc-min-version --version="3.0.0"  --proto_path=$(GOPATH)/src/github.com/gogo/protobuf:$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf:.  --micro_out=.  --gogo_out=. proto.proto)



