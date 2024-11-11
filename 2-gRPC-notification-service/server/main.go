package main

import (
	"log"
	"net"

	pb "github.com/40471/go-grpc-examples/2-gRPC-notification-service/proto"
	"github.com/40471/go-grpc-examples/2-gRPC-notification-service/server/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	notificationServer := server.NewServer()
	pb.RegisterNotificationServiceServer(s, notificationServer)

	log.Println("Server listening at", lis.Addr())

	go notificationServer.StartConsumer("news", "sports")
	go notificationServer.StartConsumer("alerts", "weather")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
