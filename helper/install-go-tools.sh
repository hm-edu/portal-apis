#!/bin/bash

install_dependencies() {    
    GOPATH="$(go env GOPATH)/bin"
    PATH="$PATH:$HOME/.local/protoc/bin:$GOPATH"
    PB_REL="https://github.com/protocolbuffers/protobuf/releases"
    VERSION="protoc-21.1-linux-x86_64.zip"
    curl -LO $PB_REL/download/v21.1/$VERSION
    mkdir -p $HOME/.local/protoc
    unzip $VERSION -d $HOME/.local/protoc
    rm $VERSION

    go install github.com/swaggo/swag/cmd/swag@latest
    
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/golang/mock/mockgen@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest    
}

go_protoc() {

    GOPATH="$(go env GOPATH)/bin"
    PATH="$HOME/.local/protoc/bin:$PATH:$GOPATH"
    # shellcheck disable=2035
    protoc \
    -I $HOME/.local/protoc/include -I. \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    *.proto
}