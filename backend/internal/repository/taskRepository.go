package repository

import "bgray/taskApi/internal/domain"

type TaskRepository interface {
	GetById(id int) (*domain.Task, error)
	GetAllTasks() ([]domain.Task, error)
	CreateTask(task domain.Task) (*domain.Task, error)
	DeleteTaskById(id int) error
	CompletedTask(id int, status bool) (*domain.Task, error)
	GetTasksByUserID(userID int) ([]domain.Task, error)
}
