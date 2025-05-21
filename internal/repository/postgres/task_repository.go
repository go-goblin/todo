package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"todo/internal/models"
	"todo/internal/repository"
)

type TaskRepository struct {
	Repository
}

func NewTaskRepository(connector *DBConnector) repository.TaskRepository {
	return &TaskRepository{Repository{pool: connector.Pool}}
}

func scanTaskRow(row pgx.Row) (*models.TaskDB, error) {
	var task models.TaskDB

	if err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdateAt,
		&task.UserID,
	); err != nil {
		return nil, err
	}

	return &task, nil
}

func scanTaskListRow(rows pgx.Rows) ([]models.TaskDB, error) {
	var items []models.TaskDB
	for rows.Next() {
		var task models.TaskDB

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdateAt,
			&task.UserID,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, task)
	}
	return items, nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context, userID int) ([]models.TaskDB, error) {
	query := `
		SELECT
		    id, 
		    title, 
		    description, 
		    status, 
		    created_at,
		    updated_at, 
		    user_id
		FROM tasks
		WHERE user_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, ErrSelect
	}
	tasks, err := scanTaskListRow(rows)
	if err != nil {
		return nil, ErrSelect
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, taskID int) (*models.TaskDB, error) {
	query := `
		SELECT
		    id, 
		    title, 
		    description, 
		    status, 
		    created_at,
		    updated_at, 
		    user_id
		FROM tasks
		WHERE id = $1;
	`
	row := r.pool.QueryRow(ctx, query, taskID)
	task, err := scanTaskRow(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskNotFound
		}
		return nil, ErrSelect
	}
	return task, nil
}

func (r *TaskRepository) CreateTask(ctx context.Context, title, description, status string, userID int) (int, error) {
	query := `
		INSERT INTO tasks(title, description, status, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	var taskID int
	err := r.pool.QueryRow(ctx, query, title, description, status, userID).Scan(&taskID)
	if err != nil {
		return 0, ErrCreateTask
	}
	return taskID, nil
}

func (r *TaskRepository) UpdateTaskByID(ctx context.Context, id int, title, description, status string) error {
	query := `
		UPDATE tasks
		SET title = $1,
			description = $2,
			status = $3
		WHERE id = $4;
	`
	tag, err := r.pool.Exec(ctx, query, title, description, status, id)
	if err != nil {
		return ErrUpdateTask
	}
	rows := tag.RowsAffected()
	if rows == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (r *TaskRepository) DeleteTaskByID(ctx context.Context, id int) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1;
	`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return ErrDeleteTask
	}
	if tag.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}
