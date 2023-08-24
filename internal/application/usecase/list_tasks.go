package usecase

import (
	"time"

	"github.com/CaioAP/go-tasks/internal/application/repository"
)

type ListTasksOutputDto struct {
	ID          string
	Name        string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ListTasksUseacse struct {
	TaskRepository repository.TaskRepository
}

func NewListTasksUsecase(taskRepository repository.TaskRepository) *ListTasksUseacse {
	return &ListTasksUseacse{TaskRepository: taskRepository}
}

func (u *ListTasksUseacse) Execute() ([]*ListTasksOutputDto, error) {
	tasks, err := u.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var tasksOutput []*ListTasksOutputDto
	for _, task := range tasks {
		tasksOutput = append(tasksOutput, &ListTasksOutputDto{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			Done:        task.Done,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return tasksOutput, nil
}
