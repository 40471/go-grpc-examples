package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"

	pb "gRPC-HelloWorld/proto"

	"google.golang.org/grpc"
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *calculatorServer) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	expression := strings.ReplaceAll(req.Expression, " ", "")
	re := regexp.MustCompile(`(\d+\.?\d*)([+\-*/])(\d+\.?\d*)`)
	matches := re.FindStringSubmatch(expression)

	if len(matches) != 4 {
		return nil, fmt.Errorf("no valid expresion")
	}

	operand1, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return nil, err
	}

	operator := matches[2]
	operand2, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return nil, err
	}

	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			return nil, fmt.Errorf("divide by 0")
		}
		result = operand1 / operand2
	default:
		return nil, fmt.Errorf("no valid operator")
	}

	return &pb.CalculateResponse{Result: result}, nil
}

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listen port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &calculatorServer{})

	log.Printf("Server port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error init server: %v", err)
	}
}
