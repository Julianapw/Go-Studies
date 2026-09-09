package main

import (
	"context"
	"fmt"
	"time"
)

func fakeAPI(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func Ex8() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := fakeAPI(ctx)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Success")
}
