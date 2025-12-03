package user

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/dto"
	"context"
	"errors"
	"fmt"
	"strings"
)

type userRepository interface {
	CreateUser(u domain.User) error
	SignIn(u string) (*domain.User, error)
}

type Service struct {
	repo   userRepository
	hasher domain.PasswordHasher
}

func NewService(repo userRepository, hasher domain.PasswordHasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (svc *Service) CreatedUser(ctx context.Context, req dto.CreateUserRequest) error {

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := user.Validate(); err != nil {
		return err
	}

	hashed, err := svc.hasher.Hash(user.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	user.Password = hashed

	lowerName := strings.ToLower(user.Username)
	user.Username = lowerName

	err = svc.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) SingIn(ctx context.Context, req dto.LoginRequest) (*domain.User, error) {
	lowerInput := strings.ToLower(req.Username)

	user, err := svc.repo.SignIn(lowerInput)
	if err != nil {
		return nil, err
	}
	if !svc.hasher.Verify(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil

}
