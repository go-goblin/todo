package repository

import (
	"context"
	"todo/internal/models"
)

type TaskRepository interface {
	GetAllTasks(ctx context.Context, userID int) ([]models.TaskDB, error)
	GetTaskByID(ctx context.Context, taskID int) (*models.TaskDB, error)
	CreateTask(ctx context.Context, title, description, status string, userID int) (int, error)
	UpdateTaskByID(ctx context.Context, id int, title, description, status string) error
	DeleteTaskByID(ctx context.Context, id int) error
}
