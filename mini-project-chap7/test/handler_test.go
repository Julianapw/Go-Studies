package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			w.WriteHeader(
				http.StatusOK,
			)
		},
	)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}
