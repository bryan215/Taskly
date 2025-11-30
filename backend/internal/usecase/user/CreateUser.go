package user

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
	"fmt"
)

type CreateUser struct {
	userRepository repository.UserRepository
	hasher         domain.PasswordHasher
}

func NewCreateUser(userRepository repository.UserRepository, hasher domain.PasswordHasher) *CreateUser {
	return &CreateUser{userRepository: userRepository, hasher: hasher}
}

func (u *CreateUser) Execute(user domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	hashed, err := u.hasher.Hash(user.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	user.Password = hashed

	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
