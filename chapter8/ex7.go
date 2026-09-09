package main

import "testing"

func ProcessNumbers(
	numbers []int,
) int {

	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func BenchmarkProcessNumbers(
	b *testing.B,
) {

	numbers := make(
		[]int,
		100000,
	)

	for i := 0; i < b.N; i++ {
		ProcessNumbers(numbers)
	}
}
