package ports

type CalculatorService interface {
	Add(a, b int) (int, error)
	Sub(a, b int) (int, error)
	Multiply(a, b int) (int, error)
	Divide(a, b int) (int, error)
}
