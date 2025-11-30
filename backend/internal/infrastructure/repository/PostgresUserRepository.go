package repository

import (
	"bgray/taskApi/internal/domain"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user domain.User) (*domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
