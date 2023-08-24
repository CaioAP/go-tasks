package repository

import (
	"database/sql"
	"time"

	"github.com/CaioAP/go-tasks/internal/domain/entity"
)

type TaskRepositoryDatabase struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepositoryDatabase {
	return &TaskRepositoryDatabase{DB: db}
}

func (r *TaskRepositoryDatabase) Create(task *entity.Task) error {
	_, err := r.DB.Exec(
		"INSERT INTO tasks (id, name, description, done, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		task.ID, task.Name, task.Description, task.Done, task.CreatedAt, task.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepositoryDatabase) FindAll() ([]*entity.Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*entity.Task
	for rows.Next() {
		var task entity.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *TaskRepositoryDatabase) FindOne(id string) (*entity.Task, error) {
	row := r.DB.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	var task *entity.Task
	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepositoryDatabase) Update(id string, task *entity.Task) error {
	_, err := r.DB.Exec(
		"UPDATE tasks SET name = ?, description = ?, done = ?, updated_at = ? WHERE id = ? RETURNING *",
		task.Name, task.Description, task.Done, time.Now(), id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepositoryDatabase) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
