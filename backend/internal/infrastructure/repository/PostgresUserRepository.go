package repository

import (
	"bgray/taskApi/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user domain.User) error {
	return r.db.Create(&user).Error
}

func (r *PostgresUserRepository) SignIn(identifier string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
