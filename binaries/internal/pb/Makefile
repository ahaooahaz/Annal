GEN_PROTOS = gen/todotask.pb.go gen/todotask.pb.validate.go
SRC_PROTOS = $(shell find . -maxdepth 1 -name "*.proto")


genpb: $(GEN_PROTOS)

$(GEN_PROTOS): $(SRC_PROTOS)
	protoc \
	-I$$GOPATH/src \
	--proto_path=. \
	--go_out=:. \
	--validate_out=lang=go:. $(SRC_PROTOS)
	protoc \
	-I$$GOPATH/src \
	--proto_path=. \
	--gotag_out=:. $(SRC_PROTOS)

clean:
	rm -rf $(GEN_PROTOS)