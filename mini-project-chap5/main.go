package main

import (
	"fmt"
)

func main() {

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://example.com",
	}

	jobs := make(chan string)
	results := make(chan Result)

	// Worker Pool
	for w := 1; w <= 3; w++ {
		go worker(jobs, results)
	}

	// Fan-Out
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Fan-In
	for i := 0; i < len(urls); i++ {

		result := <-results

		if result.Error != nil {

			fmt.Println(
				"Error fetching:",
				result.URL,
			)

			continue
		}

		fmt.Printf(
			"Fetched %s in %d ms (Status %d)\n",
			result.URL,
			result.Duration,
			result.StatusCode,
		)
	}
}
