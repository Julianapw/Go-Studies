package main

import (
	"encoding/json"
	"net/http"
)

type EchoRequest struct {
	Message string `json:"message"`
}

func Ex5() {

	http.HandleFunc("/echo", func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		var req EchoRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(
				w,
				"Invalid JSON",
				http.StatusBadRequest,
			)
			return
		}

		json.NewEncoder(w).Encode(req)
	})

	http.ListenAndServe(":8080", nil)
}
