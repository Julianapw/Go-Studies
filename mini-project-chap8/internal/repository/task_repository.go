package repository

import (
	"chap8/internal/model"
	"errors"
)

type TaskRepository struct {
	tasks []model.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: []model.Task{},
	}
}

func (r *TaskRepository) GetAll() []model.Task {
	return r.tasks
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {

	for _, task := range r.tasks {

		if task.ID == id {
			return &task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (r *TaskRepository) Create(task model.Task) {
	r.tasks = append(r.tasks, task)
}

func (r *TaskRepository) Delete(id int) error {

	for i, task := range r.tasks {

		if task.ID == id {

			r.tasks =
				append(
					r.tasks[:i],
					r.tasks[i+1:]...,
				)

			return nil
		}
	}

	return errors.New("task not found")
}
