package domain

import (
	"errors"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Tasks    []Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}
