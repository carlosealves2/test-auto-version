package service

import "testing"

func TestSubtractTwoNumbers(t *testing.T) {
	result := Sub(10, 5)
	expected := 5

	if result != expected {
		t.Errorf("Sub(10, 5) = %d; want %d", result, expected)
	}
}

func TestSubtractWithZero(t *testing.T) {
	result := Sub(0, 0)
	expected := 0

	if result != expected {
		t.Errorf("Sub(0, 0) = %d; want %d", result, expected)
	}
}

func TestSubtractNegativeNumbers(t *testing.T) {
	result := Sub(-5, -3)
	expected := -2

	if result != expected {
		t.Errorf("Sub(-5, -3) = %d; want %d", result, expected)
	}
}

func TestSubtractResultingInNegative(t *testing.T) {
	result := Sub(3, 5)
	expected := -2

	if result != expected {
		t.Errorf("Sub(3, 5) = %d; want %d", result, expected)
	}
}
