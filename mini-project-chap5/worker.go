package main

import (
	"net/http"
	"time"
)

func worker(jobs <-chan string, results chan<- Result) {

	for url := range jobs {

		start := time.Now()

		resp, err := http.Get(url)

		duration := time.Since(start)

		if err != nil {

			results <- Result{
				URL:   url,
				Error: err,
			}

			continue
		}

		results <- Result{
			URL:        url,
			StatusCode: resp.StatusCode,
			Duration:   duration.Milliseconds(),
		}

		resp.Body.Close()
	}
}
