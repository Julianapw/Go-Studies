package main

import (
	"fmt"
	"net/http"

	"./handler"
	"./service"

	"./middleware"
)

func main() {
	userService := service.NewUserService()

	userHandler := handler.UserHandler{
		Service: userService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", userHandler.Health)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	handlerWithMiddleware := middleware.Logging(middleware.Header(mux))

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", handlerWithMiddleware)
}
