package service

import "testing"

func TestTowEPlusTowEqualsFour(t *testing.T) {
	result := Sum(2, 2)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}
