package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	server := &http.Server{
		Addr: ":8080",
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	go server.ListenAndServe()

	<-ctx.Done()

	shutdownCtx, cancel :=
		context.WithTimeout(
			context.Background(),
			5*time.Second,
		)

	defer cancel()

	server.Shutdown(shutdownCtx)
}
