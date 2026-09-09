package main

import (
	"context"
	"errors"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, user User) error
}

type UserServiceImpl struct{}

func (s UserServiceImpl) CreateUser(
	ctx context.Context,
	user User,
) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Age <= 0 {
		return errors.New("age must be greater than zero")
	}

	select {
	case <-time.After(1 * time.Second):
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}
