package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Name        string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(name string, description string) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) IsDone() {
	t.Done = true
}
