package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Work completed")

	case <-ctx.Done():
		fmt.Println("Client disconnected:", ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	go doWork(ctx)

	w.Write([]byte("Request received"))
}
