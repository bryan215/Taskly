package domain

import "errors"

type Task struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	UserID    int    `json:"user_id" gorm:"not null"`
	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
	User      User   `json:"-" gorm:"foreignKey:UserID"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
