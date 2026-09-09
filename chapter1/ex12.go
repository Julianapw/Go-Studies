package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age %d: must not be negative", age)
	}
	if age < 18 {
		return fmt.Errorf("age %d is not allowed: must be 18 or older", age)
	}
	return nil
}

func Ex12() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an age: ")
	raw, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	raw = strings.TrimSpace(raw)
	age, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Println("Invalid age. Please enter a whole number (e.g., 18).")
		return
	}

	if err := ValidateAge(age); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("OK: age accepted")
}
