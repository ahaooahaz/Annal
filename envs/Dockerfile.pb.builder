FROM ubuntu:18.04

RUN apt-get update

RUN apt-get install wget git unzip make -y

RUN wget https://go.dev/dl/go1.17.13.linux-amd64.tar.gz >/dev/null 2>&1 && tar -C /usr/local -xzf go1.17.13.linux-amd64.tar.gz && rm -rf go1.17.13.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="/go"
ENV PATH="${PATH}:${GOPATH}/bin"

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protoc-3.20.3-linux-x86_64.zip >/dev/null 2>&1 \
    && unzip protoc-3.20.3-linux-x86_64.zip -d protoc \
    && mv protoc/bin/protoc /usr/local/bin \
    && mv protoc/include/google /usr/local/include \
    && rm -rf protoc-3.20.3-linux-x86_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/envoyproxy/protoc-gen-validate@latest \
    && go install github.com/srikrsna/protoc-gen-gotag@latest

RUN git clone https://github.com/googleapis/googleapis.git /go/src/github.com/googleapis/googleapis
RUN git clone https://github.com/srikrsna/protoc-gen-gotag.git /go/src/github.com/srikrsna/protoc-gen-gotag
RUN git clone https://github.com/bufbuild/protoc-gen-validate.git /go/src/github.com/envoyproxy/protoc-gen-validate