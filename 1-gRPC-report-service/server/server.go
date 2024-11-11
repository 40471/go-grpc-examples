package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/40471/go-grpc-examples/1-gRPC-report-service/proto"

	"github.com/40471/go-grpc-examples/1-gRPC-report-service/report"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedReportServiceServer
	reports map[string]*report.Report
}

func (s *server) GenerateReport(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	reportID := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + strconv.Itoa(rand.Intn(1000))
	fmt.Printf("Report Requested: Report Type: %s, ID: %s", req.ReportType, reportID)
	s.reports[reportID] = &report.Report{
		ID:        reportID,
		Type:      req.ReportType,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	go func() {
		generatedReport := report.GenerateReport(reportID, req.ReportType)
		s.reports[reportID] = generatedReport
	}()

	return &pb.ReportResponse{ReportId: reportID}, nil
}

func (s *server) StreamReportStatus(req *pb.StatusRequest, stream pb.ReportService_StreamReportStatusServer) error {
	reportID := req.ReportId
	for {
		log.Printf("Status requested for report ID: %s", reportID)
		report, exists := s.reports[reportID]
		if !exists {
			return fmt.Errorf("error report dont exists")
		}

		var reportURL string
		if report.Status == "completed" {
			reportURL = report.GetReportURL()
		}

		if err := stream.Send(&pb.StatusResponse{
			ReportId:  reportID,
			Status:    report.Status,
			ReportUrl: reportURL,
		}); err != nil {
			return err
		}

		if report.Status == "completed" {
			log.Printf("Report completed , stream close ID: %s", reportID)
			break
		}

		time.Sleep(2 * time.Second)
	}
	return nil
}

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listen port: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterReportServiceServer(grpcServer, &server{reports: make(map[string]*report.Report)})
	log.Printf("Server port %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error init server: %v", err)
	}
}
