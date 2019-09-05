## $ . utilcommand.sh

upserver() {
    for c in "go run $(find -name server)/main.go" "go run $(find -name api)/main.go"; do 
        ${c} &
    done
}

downserver() {
    ps -ef | grep server/main.go | grep -v grep | awk '{print $1}' | xargs -I{} kill -9 {}
    ps -ef | grep server/api.go  | grep -v grep | awk '{print $1}' | xargs -I{} kill -9 {}
    ps -ef | grep go-build | grep -v grep | awk '{print $1}' | xargs -I{} kill -9 {}
}

pbbuild() {
    cd $GOPATH/src/user/
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. pb/user.proto
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. pb/user.proto
}