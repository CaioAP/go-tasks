package usecase

import (
	"time"

	"github.com/CaioAP/go-tasks/internal/application/repository"
	"github.com/CaioAP/go-tasks/internal/domain/entity"
)

type UpdateTaskInputDto struct {
	Name        string
	Description string
	Done        bool
}

type UpdateTaskOutputDto struct {
	ID          string
	Name        string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateTaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func NewUpdateTaskUsecase(taskRepository repository.TaskRepository) *UpdateTaskUsecase {
	return &UpdateTaskUsecase{TaskRepository: taskRepository}
}

func (u *UpdateTaskUsecase) Execute(id string, input UpdateTaskInputDto) (*UpdateTaskOutputDto, error) {
	task := entity.NewTask(input.Name, input.Description)
	if input.Done {
		task.IsDone()
	}
	err := u.TaskRepository.Update(id, task)
	if err != nil {
		return nil, err
	}
	taskUpdated, err := u.TaskRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return &UpdateTaskOutputDto{
		ID:          taskUpdated.ID,
		Name:        taskUpdated.Name,
		Description: taskUpdated.Description,
		Done:        taskUpdated.Done,
		CreatedAt:   taskUpdated.CreatedAt,
		UpdatedAt:   taskUpdated.UpdatedAt,
	}, nil
}
