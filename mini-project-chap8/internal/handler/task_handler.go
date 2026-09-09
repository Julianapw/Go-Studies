package handler

import (
	"chap8/internal/service"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h TaskHandler) Health(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).Encode(
		map[string]string{
			"status": "UP",
		},
	)
}

func (h TaskHandler) GetTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).Encode(
		h.Service.GetTasks(),
	)
}

func (h TaskHandler) CreateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req struct {
		Title string `json:"title"`
	}

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	err = h.Service.CreateTask(
		req.Title,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(
		http.StatusCreated,
	)
}
