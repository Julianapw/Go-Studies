package main

import (
	"context"
	"fmt"
	"time"
)

func Ex11() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processing stopped")
			return

		case value := <-ch:
			fmt.Println("Received:", value)
		}
	}
}
