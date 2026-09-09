package main

import (
	"context"
	"fmt"
	"time"
)

func workerEx7(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker shutting down...")
			return

		default:
			fmt.Println("Processing job...")
			time.Sleep(1 * time.Second)
		}
	}
}

func Ex7() {
	ctx, cancel := context.WithCancel(context.Background())

	go workerEx7(ctx)

	time.Sleep(4 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}
