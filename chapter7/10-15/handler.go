package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func HealthHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(
		map[string]string{
			"status": "UP",
		},
	)
}

func UserHandler(
	service UserService,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		if r.Method != http.MethodPost {

			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)

			return
		}

		var user User

		err := json.NewDecoder(
			r.Body,
		).Decode(&user)

		if err != nil {

			http.Error(
				w,
				"invalid JSON",
				http.StatusBadRequest,
			)

			return
		}

		// Exercise 13
		ctx, cancel := context.WithTimeout(
			r.Context(),
			2*time.Second,
		)

		defer cancel()

		err = service.CreateUser(
			ctx,
			user,
		)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)

			return
		}

		w.WriteHeader(
			http.StatusCreated,
		)

		json.NewEncoder(w).Encode(
			Response{
				Message: "user created",
			},
		)
	}
}
