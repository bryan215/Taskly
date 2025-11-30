package task

import (
	"bgray/taskApi/internal/repository"
	"fmt"
)

type DeleteTaskById struct {
	taskRepository repository.TaskRepository
}

func NewDeleteTaskById(taskRepository repository.TaskRepository) *DeleteTaskById {
	return &DeleteTaskById{taskRepository: taskRepository}
}

func (d *DeleteTaskById) Execute(id int) (string, error) {
	err := d.taskRepository.DeleteTaskById(id)
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("Se ha eliminado el siguiente id: %d", id)
	return message, nil
}
