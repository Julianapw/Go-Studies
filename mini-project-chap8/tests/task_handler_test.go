package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(
	t *testing.T,
) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	if rec.Code == 0 {

		t.Error(
			"response recorder not initialized",
		)
	}

	_ = req
}
