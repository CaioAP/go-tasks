package usecase

import (
	"github.com/CaioAP/go-tasks/internal/application/repository"
)

type DeleteTaskInputDto struct {
}

type DeleteTaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func NewDeleteTaskUsecase(taskRepository repository.TaskRepository) *DeleteTaskUsecase {
	return &DeleteTaskUsecase{TaskRepository: taskRepository}
}

func (u *DeleteTaskUsecase) Execute(id string) error {
	err := u.TaskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
