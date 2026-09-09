package main

import (
	"fmt"
	"os"
	"testing"
)

func LoadFile() (string, error) {

	data, err := os.ReadFile("data.txt")

	if err != nil {
		return "", fmt.Errorf(
			"failed to read file: %w",
			err,
		)
	}

	return string(data), nil
}

func TestLoadFile(t *testing.T) {

	_, err := LoadFile()

	if err == nil {
		t.Error("expected error")
	}
}
