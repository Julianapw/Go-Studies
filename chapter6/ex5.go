package main

import (
	"context"
	"fmt"
)

func repository(ctx context.Context) {
	fmt.Println("Repository layer")
}

func service(ctx context.Context) {
	fmt.Println("Service layer")
	repository(ctx)
}

func Ex5() {
	ctx := context.Background()

	fmt.Println("Main layer")
	service(ctx)
}
