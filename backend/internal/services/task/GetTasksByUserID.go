package task

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
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
	// Devolver array vacío en lugar de error si no hay tareas
	return tasks, nil
}
