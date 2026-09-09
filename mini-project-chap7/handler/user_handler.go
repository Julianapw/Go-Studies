package handler

import (
	"encoding/json"
	"net/http"

	"../model"
)

type UserService interface {
	CreateUser(model.User) error
	GetUsers() []model.User
}

type UserHandler struct {
	Service UserService
}

func (h UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created",
	})
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.Service.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
