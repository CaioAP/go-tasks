package repository

import "github.com/CaioAP/go-tasks/internal/domain/entity"

type TaskRepository interface {
	Create(task *entity.Task) error
	FindAll() ([]*entity.Task, error)
	FindOne(id string) (*entity.Task, error)
	Update(id string, task *entity.Task) error
	Delete(id string) error
}
