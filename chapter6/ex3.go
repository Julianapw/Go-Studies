package main

import (
	"context"
	"fmt"
	"time"
)

func longTask(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func expirate() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	fmt.Println("Task started...")

	err := longTask(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
