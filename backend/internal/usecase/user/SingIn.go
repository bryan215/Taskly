package user

import (
	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/repository"
	"errors"
)

type SingIn struct {
	singInRepository repository.UserRepository
	hasher           domain.PasswordHasher
}

func NewSingIn(singInRepository repository.UserRepository, hasher domain.PasswordHasher) *SingIn {
	return &SingIn{singInRepository: singInRepository, hasher: hasher}

}

func (s *SingIn) Execute(u string, p string) (*domain.User, error) {

	user, err := s.singInRepository.SignIn(u)
	if err != nil {
		return nil, err
	}
	if !s.hasher.Verify(p, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil

}
