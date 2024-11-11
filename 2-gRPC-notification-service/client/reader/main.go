package main

import (
	"context"
	"log"
	"time"

	pb "github.com/40471/go-grpc-examples/2-gRPC-notification-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb.SubscriptionRequest{
		UserId:    "user1",
		EventType: "news",
		Filter:    "sports",
	})
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		notification, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving notification: %v", err)
			break
		}
		log.Printf("Received notification: %s", notification.Message)
	}
}
