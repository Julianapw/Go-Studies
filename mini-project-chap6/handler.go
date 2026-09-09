package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

func userHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
	)
	defer cancel()

	var response UserResponse

	var wg sync.WaitGroup

	profileCh := make(chan string, 1)
	ordersCh := make(chan string, 1)
	recommendationsCh := make(chan string, 1)

	errorCh := make(chan error, 3)

	wg.Add(3)

	go func() {
		defer wg.Done()

		data, err := fetchUserProfile(ctx)

		if err != nil {
			errorCh <- err
			return
		}

		profileCh <- data
	}()

	go func() {
		defer wg.Done()

		data, err := fetchUserOrders(ctx)

		if err != nil {
			errorCh <- err
			return
		}

		ordersCh <- data
	}()

	go func() {
		defer wg.Done()

		data, err := fetchUserRecommendations(ctx)

		if err != nil {
			errorCh <- err
			return
		}

		recommendationsCh <- data
	}()

	wg.Wait()

	close(errorCh)

	for err := range errorCh {
		response.Error = err.Error()
	}

	select {
	case response.Profile = <-profileCh:
	default:
	}

	select {
	case response.Orders = <-ordersCh:
	default:
	}

	select {
	case response.Recommendations = <-recommendationsCh:
	default:
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
