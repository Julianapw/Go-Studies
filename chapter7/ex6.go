package main

import (
	"encoding/json"
	"net/http"
)

type Echo struct {
	Message string `json:"message"`
}

func Ex6() {

	http.HandleFunc("/echo", func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		var req Echo

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(
				w,
				"Invalid JSON",
				http.StatusBadRequest,
			)
			return
		}

		if req.Message == "" {
			http.Error(
				w,
				"message is required",
				http.StatusBadRequest,
			)
			return
		}

		json.NewEncoder(w).Encode(req)
	})

	http.ListenAndServe(":8080", nil)
}
