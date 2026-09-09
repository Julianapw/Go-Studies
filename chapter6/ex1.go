package main

import (
	"context"
	"fmt"
)

func printMessage(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	fmt.Println("Hello from function!")
	_ = ctx
}

func print() {
	ctx := context.Background()
	printMessage(ctx)
}
