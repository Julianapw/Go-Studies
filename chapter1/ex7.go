package main

import "fmt"

func Ex7() {
	for i := 0; i < 21; i++ {
		fmt.Printf("%d ", i)
	}

	i := 0
	for i < 21 {
		fmt.Printf("%d ", i)
		i++
	}

	for {
		fmt.Println("infinite loop")
	}

}
