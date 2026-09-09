package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"chap8/internal/config"
	"chap8/internal/handler"
	"chap8/internal/middleware"
	"chap8/internal/repository"
	"chap8/internal/service"
)

func main() {

	cfg := config.Load()

	repo := repository.NewTaskRepository()

	service := service.NewTaskService(repo)

	handler := handler.TaskHandler{
		Service: service,
	}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/health",
		handler.Health,
	)

	mux.HandleFunc(
		"/tasks",
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			switch r.Method {

			case http.MethodGet:
				handler.GetTasks(w, r)

			case http.MethodPost:
				handler.CreateTask(w, r)

			default:
				http.Error(
					w,
					"method not allowed",
					http.StatusMethodNotAllowed,
				)
			}
		},
	)

	server := &http.Server{
		Addr: ":" + cfg.Port,
		Handler: middleware.Logging(
			mux,
		),
	}

	go func() {

		log.Printf(
			"server started on %s",
			cfg.Port,
		)

		server.ListenAndServe()
	}()

	ctx, stop :=
		signal.NotifyContext(
			context.Background(),
			os.Interrupt,
		)

	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel :=
		context.WithTimeout(
			context.Background(),
			5*time.Second,
		)

	defer cancel()

	server.Shutdown(
		shutdownCtx,
	)
}
