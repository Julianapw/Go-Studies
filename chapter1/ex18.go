package main

import "testing"

func Ex18(a int, b int) int {
	return a * b
}

func TestEx18(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{5, 3, 15},
		{2, 4, 8},
		{0, 10, 0},
	}

	for _, test := range tests {
		result := Ex18(test.a, test.b)
		if result != test.expected {
			t.Errorf("Expected%d, got %d", test.expected, result)
		}
	}
}
