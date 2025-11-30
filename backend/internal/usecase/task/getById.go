package task

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
	"errors"
)

type GetByIdTaskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewGetByIdTaskUseCase(taskRepository repository.TaskRepository) *GetByIdTaskUseCase {
	return &GetByIdTaskUseCase{taskRepository: taskRepository}
}

func (u *GetByIdTaskUseCase) Execute(id int) (*domain.Task, error) {
	task, err := u.taskRepository.GetById(id)
	if err != nil {
		return nil, errors.New("task not found")
	}
	if task == nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}
