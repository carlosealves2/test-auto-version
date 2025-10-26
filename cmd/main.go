package main

import (
	"fmt"
	"log"
	"net"

	"github.com/carlosealves2/test-auto-version/configs"
	grpcserver "github.com/carlosealves2/test-auto-version/internal/grpc"
	"github.com/carlosealves2/test-auto-version/internal/service"
	pb "github.com/carlosealves2/test-auto-version/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("iniciando gRPC server...")

	cfg := configs.NewConfigBuilder().WithEnv().Validate().Build()

	// Create calculator service
	calculatorService := service.NewCalculatorService()

	// Create gRPC server
	grpcServer := grpc.NewServer()
	calculatorServer := grpcserver.NewCalculatorServer(calculatorService)
	pb.RegisterCalculatorServer(grpcServer, calculatorServer)

	// Start listening
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server running on port %s\n", cfg.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
