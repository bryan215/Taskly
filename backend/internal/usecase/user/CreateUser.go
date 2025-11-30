package user

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
)

type CreateUser struct {
	userRepository repository.UserRepository
}

func NewCreateUser(userRepository repository.UserRepository) *CreateUser {
	return &CreateUser{userRepository: userRepository}
}

func (u *CreateUser) Execute(user domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
