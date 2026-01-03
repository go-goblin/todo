package app

import (
	"context"
	"log"
	"url-stortener/internal/app/server"
	"url-stortener/internal/dependencies"
	"url-stortener/internal/service"
)

type App struct {
	Server *server.HTTPServer
}

func New() (*App, error) {
	log.Println("🔧 Инициализация сервиса сокращения URL...")

	urlShortenerService := service.NewUrlShortenerService()
	deps := dependencies.New(urlShortenerService)
	srv := server.New(deps)

	log.Println("✅ Сервисы инициализированы")

	app := &App{
		Server: srv,
	}
	return app, nil
}

func (a *App) Run() error {
	log.Println("🚀 Запускаю HTTP сервер...")
	log.Println("🌐 Сервер будет доступен по адресу: http://localhost:8080")
	log.Println("📝 API эндпоинты:")
	log.Println("   POST /api/shorten - Создать короткую ссылку")
	log.Println("   GET  /{code}      - Перейти по короткой ссылке")
	if err := a.Server.Run(); err != nil {
		return err
	}
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}
