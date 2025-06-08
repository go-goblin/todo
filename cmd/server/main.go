package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"todo/internal/app"
	"todo/internal/config"
	"todo/internal/logger"
)

// @title To Do API
// @version 1.0
// @description API для управления задачами
// @BasePath /

// Описание схемы авторизации (добавляется в "Authorize" в Swagger UI)
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if env := os.Getenv("ENV"); env == "" || env == "local" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	logger.Init(cfg.LogLevel)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	application, err := app.New(cfg)
	if err != nil {
		panic(err)
	}

	if err := application.Run(ctx, cfg); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Get().Info(context.Background(), "остановка приложения", logrus.Fields{
				"err": err.Error(),
			})
			return
		}
		panic(err)
	}
}
