package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return

		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func run() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)

	fmt.Println("Context cancelled")
	cancel()

	time.Sleep(1 * time.Second)
}
