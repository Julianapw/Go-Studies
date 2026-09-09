package main

import (
	"context"
	"fmt"
	"time"
)

func longTaskWithCancel(ctx context.Context) {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return

		default:
			fmt.Printf("Iteration %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}
}

func Ex4() {
	ctx, cancel := context.WithCancel(context.Background())

	go longTaskWithCancel(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}
