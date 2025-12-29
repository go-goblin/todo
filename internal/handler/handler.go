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
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/{code}", h.RedirectToUrl)
		r.Post("/create", h.PostMakeShortUrl)
	})

	return r
}
