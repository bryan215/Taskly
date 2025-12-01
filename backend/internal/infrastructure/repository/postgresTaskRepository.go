package repository

import (
	"bgray/taskApi/internal/domain"
	taskRepo "bgray/taskApi/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type PostgresTaskRepository struct {
	db *gorm.DB
}

func NewPostgresTaskRepository(db *gorm.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) GetById(id int) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *PostgresTaskRepository) GetAllTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, errors.New("failed to get all tasks")
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) CreateTask(task domain.Task) (*domain.Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *PostgresTaskRepository) DeleteTaskById(id int) error {
	var task domain.Task

	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}

	if err := r.db.Delete(&task, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostgresTaskRepository) CompletedTask(id int, status bool) (*domain.Task, error) {
	var task domain.Task

	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	if err := r.db.Model(&task).Where("id = ?", id).Update("completed", status).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *PostgresTaskRepository) GetTasksByUserID(userID int) ([]domain.Task, error) {
	var tasks []domain.Task

	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

var _ taskRepo.TaskRepository = (*PostgresTaskRepository)(nil)
