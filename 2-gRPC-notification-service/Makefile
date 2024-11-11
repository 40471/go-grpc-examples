PROTOC = protoc
PROTO_DIR = proto
GO_OUT = .
PROTO_FILES = proto/notification.proto

MODULE_NAME = github.com/40471/go-grpc-examples/2-gRPC-notification-service


all: init deps generate
	docker-compose up --build
	@echo  "\033[1;32mProyecto construido\033[0m"

init:
	@echo "\033[36mInicializando el módulo Go...\033[0m"
	go mod init $(MODULE_NAME)

deps:
	@echo "\033[36mDescargando dependencias de gRPC y protobuf...\033[0m"
	go get google.golang.org/grpc
	go get google.golang.org/protobuf
	go get github.com/rabbitmq/amqp091-go

generate: deps
	@echo "\033[33mGenerando código Go a partir de archivos .proto...\033[0m"
	$(PROTOC) --go_out=$(GO_OUT) --go_opt=paths=source_relative --go-grpc_out=$(GO_OUT) --go-grpc_opt=paths=source_relative $(PROTO_FILES) 

clean:
	@echo "\033[31mLimpiando archivos generados...\033[0m"
	rm -f go.mod go.sum
	rm -f $(PROTO_DIR)/*.pb.go

re: clean all
	@echo "\033[1;34mReconstrucción completa realizada\033[0m"