package handler

import (
	"github.com/go-chi/chi/v5"
	"url-stortener/internal/dependencies"
)

type Handler struct {
	deps *dependencies.Dependencies
}

func New(deps *dependencies.Dependencies) *Handler {
	return &Handler{
		deps: deps,
	}
}

func (h *Handler) GetRouter() chi.Router {
	// TODO: Создать новый роутер chi.Router с помощью chi.NewRouter()
	// TODO: Определить маршруты в группе "/":
	//   - GET /{code}     → обработчик h.RedirectToUrl (редирект по короткому коду)
	//   - POST /create    → обработчик h.PostMakeShortUrl (создание короткой ссылки)
	// TODO: Вернуть созданный роутер
	return nil
}
