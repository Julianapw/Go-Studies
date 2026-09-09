package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("User not found")

func FindUser(id int) (string, error) {
	if id == 12 {
		return "", ErrUserNotFound
	}
	return "Juliana", nil
}

// Exercise 7
func Err() {
	result, err := FindUser(13)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("User not found")
		}
	}
	fmt.Println(result)
}

// Exercise 8
func Err2(id string) (string, error) {
	if id != "12" {
		return "", fmt.Errorf("database lookup failed: %w", ErrUserNotFound)
	}
	return "Juliana", nil
}
