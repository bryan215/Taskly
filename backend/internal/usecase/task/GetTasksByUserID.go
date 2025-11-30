package task

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
	"errors"
)

type GetTasksByUserID struct {
	taskRepository repository.TaskRepository
}

func NewGetTasksByUserID(taskRepository repository.TaskRepository) *GetTasksByUserID {
	return &GetTasksByUserID{taskRepository: taskRepository}
}

func (t *GetTasksByUserID) Execute(userID int) ([]domain.Task, error) {
	tasks, err := t.taskRepository.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found for user")
	}
	return tasks, nil
}
