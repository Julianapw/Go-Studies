package main

import (
	"context"
	"fmt"
	"time"
)

func workerShared(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return

		default:
			fmt.Printf("Worker %d running\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Ex9() {
	ctx, cancel := context.WithCancel(context.Background())

	go workerShared(ctx, 1)
	go workerShared(ctx, 2)
	go workerShared(ctx, 3)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}
