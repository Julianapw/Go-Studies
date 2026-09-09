package main

import (
	"encoding/json"
	"net/http"
)

func Ex4() {

	http.HandleFunc("/user", func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		response := map[string]string{
			"name": "Aditya",
			"role": "Developer",
		}

		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
