package task

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
)

type CreateTask struct {
	taskRepository repository.TaskRepository
}

func NewCreateTask(taskRepository repository.TaskRepository) *CreateTask {
	return &CreateTask{taskRepository: taskRepository}
}

func (t *CreateTask) Execute(task domain.Task) (*domain.Task, error) {
	if err := task.Validate(); err != nil {
		return nil, err
	}
	createdTask, err := t.taskRepository.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
