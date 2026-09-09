package main

import "fmt"

func compute() int {
	return -3
}

func Ex6() {
	if number := compute(); number == 0 {
		fmt.Printf("%d is zero.\n", number)
	} else if number > 0 {
		fmt.Printf("%d is a positive number.\n", number)
	} else {
		fmt.Printf("%d is a negative number.\n", number)
	}
}
