package usecase

import (
	"time"

	"github.com/CaioAP/go-tasks/internal/application/repository"
)

type GetTaskOutputDto struct {
	ID          string
	Name        string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetTaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func NewGetTaskUsecase(taskRepository repository.TaskRepository) *GetTaskUsecase {
	return &GetTaskUsecase{TaskRepository: taskRepository}
}

func (u *GetTaskUsecase) Execute(id string) (*GetTaskOutputDto, error) {
	task, err := u.TaskRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return &GetTaskOutputDto{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}
