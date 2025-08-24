package handler

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"todo/internal/dependencies"
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

	r.Use(h.corsMiddleware())
	r.Use(h.commonWare)
	r.Use(h.handlerLogger)

	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"),
	))

	// Обслуживание статического файла swagger.json.
	r.Get("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	r.Route("/", func(r chi.Router) {
		r.Post("/signup", h.PostSignIn)
		r.Post("/login", h.PostLogin)
	})

	r.With(h.jwtAuthMiddleware).Group(func(r chi.Router) {
		r.Get("/tasks", h.GetAllTasks)
		r.Get("/tasks/{id}", h.GetTaskByID)
		r.Post("/tasks", h.PostCreateTask)
		r.Put("/tasks/{id}", h.PutUpdateTask)
		r.Delete("/tasks/{id}", h.DelDeleteTask)
	})
	return r
}
