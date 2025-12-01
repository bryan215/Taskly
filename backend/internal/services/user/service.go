package user

import (
	"bgray/taskApi/internal/domain"
	"errors"
	"fmt"
)

type userRepository interface {
	CreateUser(u domain.User) (*domain.User, error)
	SignIn(u string) (*domain.User, error)
}

type Service struct {
	repo   userRepository
	hasher domain.PasswordHasher
}

func NewService(repo userRepository, hasher domain.PasswordHasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (svc *Service) CreatedUser(user domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	hashed, err := svc.hasher.Hash(user.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	user.Password = hashed

	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := svc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (svc *Service) SingIn(u string, p string) (*domain.User, error) {

	user, err := svc.repo.SignIn(u)
	if err != nil {
		return nil, err
	}
	if !svc.hasher.Verify(p, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil

}
