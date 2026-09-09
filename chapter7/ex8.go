package main

import "net/http"

func HeaderMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		w.Header().Set(
			"X-App-Version",
			"1.0",
		)

		next.ServeHTTP(w, r)
	})
}
