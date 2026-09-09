package service

import (
	"errors"

	"../model"
)

type UserService interface {
	CreateUser(user model.User) error
	GetUsers() []model.User
}

type UserServiceImpl struct {
	users []model.User
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{
		users: []model.User{},
	}
}

func (s *UserServiceImpl) CreateUser(user model.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Age <= 0 {
		return errors.New("age must be greater than zero")
	}

	s.users = append(s.users, user)

	return nil
}

func (s *UserServiceImpl) GetUsers() []model.User {
	return s.users
}
