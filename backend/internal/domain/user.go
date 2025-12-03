package domain

import (
	"errors"
	"regexp"
)

const (
	RegexEmail = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
)

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if !regexp.MustCompile(RegexEmail).MatchString(u.Email) {
		return errors.New("invalid email")
	}
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password, hash string) bool
}

type TokenGenerator interface {
	GenerateToken(user User) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
}

type TokenClaims struct {
	UserID   int
	Username string
}
