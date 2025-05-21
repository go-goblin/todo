package handler

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
}
