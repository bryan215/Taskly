package task

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
	"errors"
)

type GetAllTasksUseCase struct {
	taskRepository repository.TaskRepository
}

func NewGetAllTasksUseCase(taskRepository repository.TaskRepository) *GetAllTasksUseCase {
	return &GetAllTasksUseCase{taskRepository: taskRepository}
}

func (u *GetAllTasksUseCase) Execute() ([]domain.Task, error) {
	tasks, err := u.taskRepository.GetAllTasks()
	if err != nil {
		return nil, errors.New("failed to get all tasks")
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}
