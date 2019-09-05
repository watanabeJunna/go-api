FROM golang:1.12.9-alpine3.10

RUN set -eux; \
    apk add --no-cache git curl protobuf \
    && go get -u google.golang.org/grpc \
    && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    && go get -u github.com/golang/protobuf/protoc-gen-go \
    && go get -u go.mongodb.org/mongo-driver/bson \
    && go get -u go.mongodb.org/mongo-driver/mongo \
    && go get -u go.mongodb.org/mongo-driver/mongo/options \
    && curl -L -o /tmp/protobuf.tar.gz https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip \
    && unzip /tmp/protobuf.tar.gz -d /tmp/. \
    && mv /tmp/include /usr/local/. \
    && rm -rf /tmp/* \
    && mkdir /go/src/user

WORKDIR /go/src/user