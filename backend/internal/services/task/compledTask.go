package task

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
)

type CompletedTask struct {
	taskRepository repository.TaskRepository
}

func NewCompletedTask(taskRepository repository.TaskRepository) *CompletedTask {
	return &CompletedTask{taskRepository: taskRepository}
}

func (t *CompletedTask) Execute(id int, status bool) (*domain.Task, error) {
	task, err := t.taskRepository.CompletedTask(id, status)
	if err != nil {
		return nil, err
	}

	return task, nil
}
