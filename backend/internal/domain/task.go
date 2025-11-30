package domain

import "errors"

type Task struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
