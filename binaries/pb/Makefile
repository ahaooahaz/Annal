GEN_PROTOS = gen/todotask.pb.go \
				gen/todotask.pb.validate.go \
				gen/serve.pb.go \
				gen/serve_grpc.pb.go \
				gen/serve.pb.validate.go
SRC_PROTOS = $(shell find . -maxdepth 1 -name "*.proto")

genpb: $(GEN_PROTOS)

$(GEN_PROTOS): $(SRC_PROTOS)
	protoc \
	-I$$GOPATH/src \
	-I$$GOPATH/src/github.com/googleapis/googleapis \
	--proto_path=. \
	--go-grpc_out=:. \
	--go_out=:. \
	--validate_out=lang=go:. $(SRC_PROTOS)
	protoc \
	-I$$GOPATH/src \
	-I$$GOPATH/src/github.com/googleapis/googleapis \
	--go-grpc_out=:. \
	--proto_path=. \
	--gotag_out=:. $(SRC_PROTOS)

clean:
	rm -rf $(GEN_PROTOS)