package service

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("Expected 12, but got %d", result)
	}
}

func TestMultiplyByZero(t *testing.T) {
	result := Multiply(5, 0)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}
