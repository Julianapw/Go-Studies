package service

import (
	"chap8/internal/model"
	"chap8/internal/repository"
	"errors"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(
	repo *repository.TaskRepository,
) *TaskService {

	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetTasks() []model.Task {
	return s.repo.GetAll()
}

func (s *TaskService) CreateTask(
	title string,
) error {

	if title == "" {
		return errors.New("title is required")
	}

	task := model.Task{
		ID:    len(s.repo.GetAll()) + 1,
		Title: title,
	}

	s.repo.Create(task)

	return nil
}

func (s *TaskService) DeleteTask(
	id int,
) error {

	return s.repo.Delete(id)
}
