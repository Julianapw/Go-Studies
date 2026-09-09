package main

import (
	"fmt"
	"net/http"
)

func Ex2() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(
				w,
				"Method Not Allowed",
				http.StatusMethodNotAllowed,
			)
			return
		}

		fmt.Fprintln(w, "Hello, Go HTTP!")
	})

	http.ListenAndServe(":8080", nil)
}
