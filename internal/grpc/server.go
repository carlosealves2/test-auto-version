package grpc

import (
	"context"

	"github.com/carlosealves2/test-auto-version/internal/ports"
	pb "github.com/carlosealves2/test-auto-version/proto"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
	calculatorService ports.CalculatorService
}

func NewCalculatorServer(calculatorService ports.CalculatorService) *CalculatorServer {
	return &CalculatorServer{
		calculatorService: calculatorService,
	}
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.calculatorService.Add(int(req.A), int(req.B))
	if err != nil {
		return &pb.OperationResponse{
			Result: 0,
			Error:  err.Error(),
		}, nil
	}

	return &pb.OperationResponse{
		Result: int32(result),
		Error:  "",
	}, nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.calculatorService.Sub(int(req.A), int(req.B))
	if err != nil {
		return &pb.OperationResponse{
			Result: 0,
			Error:  err.Error(),
		}, nil
	}

	return &pb.OperationResponse{
		Result: int32(result),
		Error:  "",
	}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.calculatorService.Multiply(int(req.A), int(req.B))
	if err != nil {
		return &pb.OperationResponse{
			Result: 0,
			Error:  err.Error(),
		}, nil
	}

	return &pb.OperationResponse{
		Result: int32(result),
		Error:  "",
	}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.calculatorService.Divide(int(req.A), int(req.B))
	if err != nil {
		return &pb.OperationResponse{
			Result: 0,
			Error:  err.Error(),
		}, nil
	}

	return &pb.OperationResponse{
		Result: int32(result),
		Error:  "",
	}, nil
}
