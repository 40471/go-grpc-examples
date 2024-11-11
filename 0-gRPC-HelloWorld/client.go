package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	pb "gRPC-HelloWorld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	// Prueba de operación con el nuevo método Calculate
	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	calcRes, err := client.Calculate(context.Background(), &pb.CalculateRequest{Expression: expr})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Result '%s' = %v", expr, calcRes.Result)
}
