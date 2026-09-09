package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type Storage interface {
	Save(data string) error
	Get(id string) (string, error)
}

type InMemoryStorage struct {
	data map[string]string
}

// Exercise 10
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]string),
	}
}

func (s *InMemoryStorage) Save(data string) error {
	s.data[data] = data
	return nil
}

func (s *InMemoryStorage) Get(id string) (string, error) {
	val, ok := s.data[id]
	if !ok {
		return "", fmt.Errorf("data not found")
	}
	return val, nil
}

// Exercise 11
type UserService struct {
	storage Storage
}

func NewUserService(storage Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (u *UserService) CreateUser(name string) error {
	return u.storage.Save(name)
}

// Exercise 12
func (u *UserService) ErrUser(name string) error {
	if err := u.storage.Save(name); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// Exercise 13
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// Exercise 14
func (u *UserService) ErrCreateUser(name string) error {
	if name == "" {
		return &ValidationError{
			Field:   "name",
			Message: "name cannot be empty",
		}
	}
	if err := u.storage.Save(name); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// Exercise 15
func (u *UserService) UserCustomError(name string) error {
	err := u.ErrCreateUser("")

	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("Validation error on field %s: %s\n", ve.Field, ve.Message)
	} else if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
	} else {
		fmt.Println("User created successfully")
	}
	return err
}

type MockStorage struct {
	ShouldFail bool
}

func (m *MockStorage) Save(data string) error {
	if m.ShouldFail {
		return fmt.Errorf("mock storage failure")
	}
	return nil
}

func (m *MockStorage) Get(id string) (string, error) {
	return "mock", nil
}

// Exercise 16TestCreateUserValidationError
func TestCreateUserSuccess(t *testing.T) {
	mock := &MockStorage{ShouldFail: false}
	service := NewUserService(mock)

	err := service.ErrCreateUser("Alice")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

// Exercise 18
func TestUserFailure(t *testing.T) {
	mock := &MockStorage{ShouldFail: true}
	service := NewUserService(mock)

	err := service.ErrCreateUser("Juliana")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "mock storage failure") {
		t.Fatalf("expected error to contain 'mock storage failure', got %v", err)
	}
}

// Excercise 19
func TestCreateUserValidationError(t *testing.T) {
	mock := &MockStorage{}
	service := NewUserService(mock)

	err := service.ErrCreateUser("")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected error to be of type ValidationError, got %v", err)
	}

	if ve.Field != "name" {
		t.Fatalf("expected validation error on field 'name', got %s", ve.Field)
	}
}
