package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"todo/internal/logger"
	"todo/internal/mapper"
	"todo/internal/models"
	"todo/internal/repository"
	"todo/internal/repository/postgres"
)

type TaskService struct {
	TaskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return TaskService{
		TaskRepository: taskRepository,
	}
}

func (s *TaskService) GetAllTasks(ctx context.Context, userId int) (*models.TaskListResponse, error) {
	logger.Get().Info(ctx, "Получение всех задач", nil)

	tasks, err := s.TaskRepository.GetAllTasks(ctx, userId)
	if err != nil {
		logger.Get().Error(ctx, "Ошибка при получении задач", logrus.Fields{
			"error": err.Error(),
		})
		return nil, err
	}

	taskDTOs := make([]models.TaskDTO, len(tasks))
	for i, e := range tasks {
		taskDTOs[i] = mapper.MapTaskDTOFromTaskDb(e)
	}

	return &models.TaskListResponse{Data: &taskDTOs}, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, userID, taskID int) (*models.TaskResponse, error) {
	task, err := s.TaskRepository.GetTaskByID(ctx, taskID)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskNotFound) {
			logger.Get().Info(ctx, "Задача не найдена", logrus.Fields{
				"task_id": taskID,
			})
			return nil, ErrTaskNotFound
		}
		logger.Get().Error(ctx, "Ошибка при получении задачи", logrus.Fields{
			"task_id": taskID,
			"error":   err.Error(),
		})
		return nil, err
	}

	if task.UserID != userID {
		logger.Get().Info(ctx, "Попытка доступа к чужой задаче", logrus.Fields{
			"task_id":  taskID,
			"owner_id": task.UserID,
		})
		return nil, ErrTaskForbidden
	}

	taskDTO := mapper.MapTaskDTOFromTaskDb(*task)
	return &models.TaskResponse{Data: &taskDTO}, nil
}

func (s *TaskService) CreateTask(ctx context.Context, input *CreateTask) (*models.CreateTaskResponse, error) {
	logger.Get().Info(ctx, "Создание новой задачи", logrus.Fields{
		"title":       input.Title,
		"description": input.Description,
		"status":      input.Status,
	})

	taskID, err := s.TaskRepository.CreateTask(ctx, input.Title, input.Description, input.Status, input.UserID)
	if err != nil {
		logger.Get().Error(ctx, "Ошибка при создании задачи", logrus.Fields{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Get().Info(ctx, "Задача успешно создана", logrus.Fields{
		"task_id": taskID,
	})

	return &models.CreateTaskResponse{Data: &models.CreateTaskData{TaskID: taskID}}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, input *UpdateTask) (*models.OperationResultResponse, error) {
	task, err := s.TaskRepository.GetTaskByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskNotFound) {
			logger.Get().Info(ctx, "Задача для обновления не найдена", logrus.Fields{
				"task_id": input.ID,
			})
			return nil, ErrTaskNotFound
		}
		logger.Get().Error(ctx, "Ошибка при получении задачи перед обновлением", logrus.Fields{
			"task_id": input.ID,
			"error":   err.Error(),
		})
		return nil, err
	}

	if task.UserID != input.UserID {
		logger.Get().Info(ctx, "Попытка обновить чужую задачу", logrus.Fields{
			"task_id":  input.ID,
			"owner_id": task.UserID,
		})
		return nil, ErrTaskForbidden
	}

	err = s.TaskRepository.UpdateTaskByID(ctx, input.ID, input.Title, input.Description, input.Status)
	if err != nil {
		logger.Get().Error(ctx, "Ошибка при обновлении задачи", logrus.Fields{
			"task_id": input.ID,
			"error":   err.Error(),
		})
		return nil, err
	}

	return &models.OperationResultResponse{Data: &models.OperationResultData{Success: true}}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, userID, taskID int) (*models.OperationResultResponse, error) {
	task, err := s.TaskRepository.GetTaskByID(ctx, taskID)
	if err != nil {
		if errors.Is(err, postgres.ErrTaskNotFound) {
			logger.Get().Info(ctx, "Задача для удаления не найдена", logrus.Fields{
				"task_id": taskID,
			})
			return nil, ErrTaskNotFound
		}
		logger.Get().Error(ctx, "Ошибка при получении задачи перед удалением", logrus.Fields{
			"task_id": taskID,
			"error":   err.Error(),
		})
		return nil, err
	}

	if task.UserID != userID {
		logger.Get().Info(ctx, "Попытка удалить чужую задачу", logrus.Fields{
			"task_id":  taskID,
			"owner_id": task.UserID,
		})
		return nil, ErrTaskForbidden
	}

	err = s.TaskRepository.DeleteTaskByID(ctx, taskID)
	if err != nil {
		logger.Get().Error(ctx, "Ошибка при удалении задачи", logrus.Fields{
			"task_id": taskID,
			"error":   err.Error(),
		})
		return nil, err
	}

	logger.Get().Info(ctx, "Задача успешно удалена", logrus.Fields{
		"task_id": taskID,
	})

	return &models.OperationResultResponse{Data: &models.OperationResultData{Success: true}}, nil
}
