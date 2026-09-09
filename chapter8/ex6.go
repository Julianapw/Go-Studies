package main

import (
	"errors"
	"testing"
)

func Divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func TestDivide(t *testing.T) {

	tests := []struct {
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true},
	}

	for _, tt := range tests {

		got, err := Divide(tt.a, tt.b)

		if tt.wantErr && err == nil {
			t.Error("expected error")
		}

		if got != tt.want {
			t.Errorf(
				"expected %d got %d",
				tt.want,
				got,
			)
		}
	}
}
