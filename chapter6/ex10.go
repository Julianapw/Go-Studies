package main

import (
	"context"
	"fmt"
	"time"
)

func safeWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cleaning up worker")
			return

		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Ex10() {
	ctx, cancel := context.WithCancel(context.Background())

	go safeWorker(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}
