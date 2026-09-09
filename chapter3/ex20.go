package main

import (
	"errors"
	"fmt"
	"testing"
)

// Interface
type Notifier interface {
	Notify(message string) error
}

// Implementations
type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) error {
	return nil
}

type SMSNotifier struct{}

func (s SMSNotifier) Notify(message string) error {
	return nil
}

// Service
type NotificationService struct {
	notifier Notifier
}

func NewNotificationService(n Notifier) *NotificationService {
	return &NotificationService{notifier: n}
}

func (s *NotificationService) SendNotification(message string) error {
	if err := s.notifier.Notify(message); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	return nil
}

// Tests
type MockNotifier struct {
	ShouldFail bool
}

func (m MockNotifier) Notify(message string) error {
	if m.ShouldFail {
		return errors.New("notification failed")
	}
	return nil
}

//Test Example
func TestNotifyFailure(t *testing.T) {
	mock := MockNotifier{ShouldFail: true}
	service := NewNotificationService(mock)

	err := service.SendNotification("Hello")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}