


go get -u github.com/golang/protobuf/protoc-gen-go

go get -u google.golang.org/grpc

protoc --go_out=plugins=grpc:. *.proto


go get -u github.com/grpc/grpc-go/examples/helloworld/greeter_client

go get -u google.golang.org/grpc/examples/helloworld/greeter_client
go get -u google.golang.org/grpc/examples/helloworld/greeter_server