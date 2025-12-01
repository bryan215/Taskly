package task

import (
	"bgray/taskApi/internal/domain"
)

type taskRepository interface {
	GetById(id int) (*domain.Task, error)
	GetAllTasks() ([]domain.Task, error)
	CreateTask(task domain.Task) (*domain.Task, error)
	DeleteTaskById(id int) error
	CompletedTask(id int, status bool) (*domain.Task, error)
	GetTasksByUserID(userID int) ([]domain.Task, error)
}

type Service struct {
	repo taskRepository
}

func NewService(repo taskRepository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) CreateTask(task domain.Task) (*domain.Task, error) {
	if err := task.Validate(); err != nil {
		return nil, err
	}
	createdTask, err := svc.repo.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
