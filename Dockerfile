FROM golang:1.6-alpine

RUN set -eux; \
    && apk add --no-cache git curl protobuf

RUN mkdir -p authenticationservice/src/genproto \
    && cd authenticationservice

COPY authencation.proto .

RUN protoc -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:./src/genproto \
    authentication.proto \
    && protoc -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:./src/genproto \
    authentication.proto

COPY src/main.go src/main.go