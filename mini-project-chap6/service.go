package main

import (
	"context"
	"time"
)

func fetchUserProfile(ctx context.Context) (string, error) {
	select {
	case <-time.After(1 * time.Second):
		return "Juliana Profile", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fetchUserOrders(ctx context.Context) (string, error) {
	select {
	case <-time.After(2 * time.Second):
		return "5 Orders", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fetchUserRecommendations(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second):
		return "Go Course", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}
