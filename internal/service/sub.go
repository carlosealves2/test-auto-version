package service

func Sub(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	return a - b
}
