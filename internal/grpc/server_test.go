package grpc

import (
	"context"
	"testing"

	"github.com/carlosealves2/test-auto-version/internal/service"
	pb "github.com/carlosealves2/test-auto-version/proto"
)

func TestCalculatorServer_Add(t *testing.T) {
	calculatorService := service.NewCalculatorService()
	server := NewCalculatorServer(calculatorService)

	req := &pb.OperationRequest{A: 5, B: 3}
	resp, err := server.Add(context.Background(), req)

	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	if resp.Result != 8 {
		t.Errorf("Expected 8, got %d", resp.Result)
	}

	if resp.Error != "" {
		t.Errorf("Expected no error, got %s", resp.Error)
	}
}

func TestCalculatorServer_Subtract(t *testing.T) {
	calculatorService := service.NewCalculatorService()
	server := NewCalculatorServer(calculatorService)

	req := &pb.OperationRequest{A: 10, B: 4}
	resp, err := server.Subtract(context.Background(), req)

	if err != nil {
		t.Fatalf("Subtract failed: %v", err)
	}

	if resp.Result != 6 {
		t.Errorf("Expected 6, got %d", resp.Result)
	}

	if resp.Error != "" {
		t.Errorf("Expected no error, got %s", resp.Error)
	}
}

func TestCalculatorServer_Multiply(t *testing.T) {
	calculatorService := service.NewCalculatorService()
	server := NewCalculatorServer(calculatorService)

	req := &pb.OperationRequest{A: 3, B: 4}
	resp, err := server.Multiply(context.Background(), req)

	if err != nil {
		t.Fatalf("Multiply failed: %v", err)
	}

	if resp.Result != 12 {
		t.Errorf("Expected 12, got %d", resp.Result)
	}

	if resp.Error != "" {
		t.Errorf("Expected no error, got %s", resp.Error)
	}
}

func TestCalculatorServer_Divide(t *testing.T) {
	calculatorService := service.NewCalculatorService()
	server := NewCalculatorServer(calculatorService)

	req := &pb.OperationRequest{A: 20, B: 4}
	resp, err := server.Divide(context.Background(), req)

	if err != nil {
		t.Fatalf("Divide failed: %v", err)
	}

	if resp.Result != 5 {
		t.Errorf("Expected 5, got %d", resp.Result)
	}

	if resp.Error != "" {
		t.Errorf("Expected no error, got %s", resp.Error)
	}
}

func TestCalculatorServer_DivideByZero(t *testing.T) {
	calculatorService := service.NewCalculatorService()
	server := NewCalculatorServer(calculatorService)

	req := &pb.OperationRequest{A: 10, B: 0}
	resp, err := server.Divide(context.Background(), req)

	if err != nil {
		t.Fatalf("Divide failed: %v", err)
	}

	if resp.Error == "" {
		t.Error("Expected error for division by zero, got none")
	}
}
