package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	service := UserServiceImpl{}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/health",
		HealthHandler,
	)

	mux.Handle(
		"/users",
		LoggingMiddleware(
			UserHandler(service),
		),
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {

		log.Println(
			"Server running on :8080",
		)

		err := server.ListenAndServe()

		if err != nil &&
			err != http.ErrServerClosed {

			log.Fatal(err)
		}
	}()

	// Exercise 14
	stop := make(
		chan os.Signal,
		1,
	)

	signal.Notify(
		stop,
		os.Interrupt,
	)

	<-stop

	log.Println(
		"Shutting down server...",
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	server.Shutdown(ctx)

	log.Println(
		"Server stopped",
	)
}
