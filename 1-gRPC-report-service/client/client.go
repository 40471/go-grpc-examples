package main

import (
	"context"
	"log"
	"time"

	pb "github.com/40471/go-grpc-examples/1-gRPC-report-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReportServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GenerateReport(ctx, &pb.ReportRequest{ReportType: "summary"})
	if err != nil {
		log.Fatalf("Error request Report: %v", err)
	}
	log.Printf("Report ID: %s", r.ReportId)

	stream, err := c.StreamReportStatus(context.Background(), &pb.StatusRequest{ReportId: r.ReportId})
	if err != nil {
		log.Fatalf("Error getting stream status: %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Printf("Error getting Status Update: %v", err)
				done <- true
				return
			}

			log.Printf("Report Status: %s", res.Status)
			if res.Status == "completed" {
				log.Printf("Eeport Url: %s", res.ReportUrl)
				done <- true
				return
			}
		}
	}()

	<-done
	log.Println("End stream")
}
