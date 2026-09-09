package main

import (
	"fmt"
	"strings"
)

func validateInput(name string, age int) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("name must not be empty")
	}
	if age < 18 {
		return fmt.Errorf("age %d is not allowed: must be 18 or older", age)
	}
	return nil
}

func Ex13(name string, age int) {
	if err := validateInput(name, age); err != nil {
		fmt.Println("Invalid input:", err) // handled here
		return
	}
	fmt.Println("All good, continue...")
}

// Why is enforced formating beneficial in teams?
//The code becomes easier to read
//Pull Requests stay cleaner (fewer changes that are only formatting)
//It’s easier for anyone to join the project and understand the style
//The team spends less time reviewing formatting and more time checking if the code is actually correct

//Why does Go treat unused imports as errors?
//Prevents bugs: an unused import usually indicates dead code or forgotten logic. The compiler warns you early.
//Keeps the code clean: less visual “noise” and more focus on what matters.
//Faster and more predictable builds: compiling only what is actually used avoids unnecessary work.
//Consistency and discipline: reinforces Go’s philosophy of being explicit and simple, without extra clutter.
//Better maintainability: PRs and diffs stay smaller and clearer (no leftover or unused imports).
