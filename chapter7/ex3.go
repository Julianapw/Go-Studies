package main

import (
	"fmt"
	"net/http"
)

type CustomHandler struct{}

func (c CustomHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(w, "Custom Handler")
}

func Ex3() {
	http.Handle("/custom", CustomHandler{})

	http.ListenAndServe(":8080", nil)
}
