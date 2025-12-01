package repository

import "bgray/taskApi/internal/domain"

type UserRepository interface {
	CreateUser(u domain.User) (*domain.User, error)
	SignIn(u string) (*domain.User, error)
}
