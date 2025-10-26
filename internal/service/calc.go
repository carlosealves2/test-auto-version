package service

import "fmt"

var (
	ErrDivideByZero = fmt.Errorf("cannot divide by zero")
)

type ImplCalculatorService struct{}

func NewCalculatorService() *ImplCalculatorService {
	return &ImplCalculatorService{}
}

func (s *ImplCalculatorService) Add(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, nil
	}
	return a + b, nil
}

func (s *ImplCalculatorService) Multiply(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}
	return a * b, nil
}

func (s *ImplCalculatorService) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func (s *ImplCalculatorService) Sub(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, nil
	}
	return a - b, nil
}
