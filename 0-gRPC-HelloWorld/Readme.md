
### **gRPC HelloWord**

- **Description**: Basic client/server calculator

go mod init gRPC-HelloWorld

go get google.golang.org/grpc

go get google.golang.org/protobuf

protoc --go_out=. --go-grpc_out=. proto/calculator.proto
