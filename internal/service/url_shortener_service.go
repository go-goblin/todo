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

type TaskService struct{}

func NewTaskService() TaskService {
	return TaskService{}
}

func (s *TaskService) CreateTask(input *CreateShortUrl) (*models.ShortUrlResultResponse, error) {

	response := &models.ShortUrlResultResponse{
		Data: &models.ShortUrlResultData{
			ShortUrl: "",
		},
	}

	return response, nil
}
