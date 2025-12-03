package task

import (
	"bgray/taskApi/internal/domain"
	"fmt"
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

func (svc *Service) CompletedTask(id int, status bool) (*domain.Task, error) {
	task, err := svc.repo.CompletedTask(id, status)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (svc *Service) DeleteTaskById(id int) (string, error) {
	err := svc.repo.DeleteTaskById(id)
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("Se ha eliminado el siguiente id: %d", id)
	return message, nil
}

func (svc *Service) GetTasksByUserID(userID int) ([]domain.Task, error) {
	tasks, err := svc.repo.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (svc *Service) GetTaskById(id int) (*domain.Task, error) {
	task, err := svc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
