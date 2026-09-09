package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ex8() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a day of the week: ")
	raw, _ := reader.ReadString('\n')
	day := strings.ToLower(strings.TrimSpace(raw))

	switch {
	case day == "monday" || day == "tuesday" || day == "wednesday" ||
		day == "thursday" || day == "friday":
		fmt.Println("Weekday")
	case day == "saturday" || day == "sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}
