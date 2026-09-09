package main

func Ex11(a int, b int, op func(int, int) int) int {
	return op(a, b)
}
