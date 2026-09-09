package main

import "testing"

func Ex16(a int, b int) int {
	return a * b
}

func TestEx16(t *testing.T) {
	result := Ex16(5, 3)
	if result != 15 {
		t.Errorf("Expected 15, but got %d", result)
	}
}
