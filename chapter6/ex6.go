package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Process completed")

	case <-ctx.Done():
		fmt.Println("Process cancelled:", ctx.Err())
	}
}

func executeWithTimeout(parent context.Context) {
	ctx, cancel := context.WithTimeout(parent, 2*time.Second)
	defer cancel()

	process(ctx)
}

func Ex6() {
	executeWithTimeout(context.Background())
}
