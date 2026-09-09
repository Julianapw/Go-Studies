package main

import "fmt"

func Ex10(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Why does the function return an error instead of throwing an exception?
// Go does not have exceptions. Instead, it uses error values to indicate that an error has occurred. This allows for more explicit error handling and makes it easier to understand the flow of the program. By returning an error value, the caller can decide how to handle the error, whether it's logging it, returning it up the call stack, or taking some other action.
