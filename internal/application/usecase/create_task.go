package usecase

import (
	"time"

	"github.com/CaioAP/go-tasks/internal/application/repository"
	"github.com/CaioAP/go-tasks/internal/domain/entity"
)

type CreateTaskInputDto struct {
	Name        string
	Description string
}

type CreateTaskOutputDto struct {
	ID          string
	Name        string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateTaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func NewCreateTaskUsecase(taskRepository repository.TaskRepository) *CreateTaskUsecase {
	return &CreateTaskUsecase{TaskRepository: taskRepository}
}

func (u *CreateTaskUsecase) Execute(input CreateTaskInputDto) (*CreateTaskOutputDto, error) {
	task := entity.NewTask(input.Name, input.Description)
	err := u.TaskRepository.Create(task)
	if err != nil {
		return nil, err
	}
	return &CreateTaskOutputDto{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}
