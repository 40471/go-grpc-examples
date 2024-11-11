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
	// Conectar al servidor gRPC
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	// Usa un contexto sin límite de tiempo para la suscripción
	subscriptionCtx := context.Background()

	// Suscribirse a notificaciones
	go func() {
		stream, err := client.Subscribe(subscriptionCtx, &pb.SubscriptionRequest{
			UserId:    "user1",
			EventType: "news",
			Filter:    "sports",
		})
		if err != nil {
			log.Fatalf("could not subscribe: %v", err)
		}

		log.Println("Subscribed to notifications: EventType='news', Filter='sports'")

		for {
			notification, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving notification: %v", err)
				break
			}
			log.Printf("Received notification: %s", notification.Message)
		}
	}()

	// Contexto separado con límite de tiempo solo para PublishEvent
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Esperar un momento y luego enviar un evento usando PublishEvent
	time.Sleep(5 * time.Second)
	log.Println("Publishing event from client: EventType='news', Filter='sports'")
	_, err = client.PublishEvent(ctx, &pb.PublishEventRequest{
		EventType: "news",
		Filter:    "sports",
		Message:   "Client-initiated event: New update on sports news!",
	})
	if err != nil {
		log.Fatalf("could not publish event: %v", err)
	}

	log.Println("Event published successfully")
	time.Sleep(time.Second * 20) // Mantiene el cliente activo para recibir notificaciones
}
