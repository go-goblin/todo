package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"todo/internal/auth"
	"todo/internal/errcodes"
	"todo/internal/httpx"
	"todo/internal/service"
)

// GetAllTasks Возвращает все задачи пользователя
// @Summary Получить задачи
// @Tags "to do"
// @Success 200 {object} models.TaskListResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /tasks [get]
func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("пользователь не авторизован"))
		return
	}
	response, err := h.deps.TaskService.GetAllTasks(ctx, userID)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}

// GetTaskByID Возвращает задачу пользователя по идентификатору
// @Summary Получить задачи по идентификатору
// @Tags "to do"
// @Success 200 {object} models.TaskResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /tasks/{id} [get]
func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("пользователь не авторизован"))
		return
	}
	taskIDStr := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusBadRequest, errors.New("неверный формат идентификатора задачи"))
		return
	}
	response, err := h.deps.TaskService.GetTaskByID(ctx, userID, taskID)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			errcodes.SendErrorJSON(w, r, http.StatusNotFound, err)
			return
		}
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}

// PostCreateTask Создать задачу
// @Summary Создает задачу для пользователя
// @Tags "to do"
// @Param request body handler.CreateTaskRequest true "Reserve Request Body"
// @Success 200 {object} models.CreateTaskResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /tasks [post]
func (h *Handler) PostCreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("пользователь не авторизован"))
		return
	}
	var request CreateTaskRequest
	if err := httpx.DecodeAndValidateBody(w, r, &request); err != nil {
		return
	}
	input := &service.CreateTask{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		UserID:      userID,
	}
	response, err := h.deps.TaskService.CreateTask(ctx, input)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}

// PutUpdateTask Обновить задачу по идентификатору
// @Summary Обновляет задачу для пользователя
// @Tags to do
// @Param id path int true "ID задачи"
// @Param request body handler.UpdateTaskRequest true "Reserve Request Body"
// @Success 200 {object} models.OperationResultResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /tasks/{id} [put]
func (h *Handler) PutUpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("пользователь не авторизован"))
		return
	}
	taskIDStr := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusBadRequest, errors.New("неверный формат идентификатора задачи"))
		return
	}
	var request UpdateTaskRequest
	if err := httpx.DecodeAndValidateBody(w, r, &request); err != nil {
		return
	}
	input := &service.UpdateTask{
		ID:          taskID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		UserID:      userID,
	}
	response, err := h.deps.TaskService.UpdateTask(ctx, input)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			errcodes.SendErrorJSON(w, r, http.StatusNotFound, err)
			return
		}
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}

// DelDeleteTask Удалить задачу по идентификатору
// @Summary Удаляет задачу для пользователя
// @Tags "to do"
// @Success 200 {object} models.OperationResultResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /tasks/{id} [put]
func (h *Handler) DelDeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("пользователь не авторизован"))
		return
	}
	taskIDStr := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusBadRequest, errors.New("неверный формат идентификатора задачи"))
		return
	}
	response, err := h.deps.TaskService.DeleteTask(ctx, userID, taskID)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			errcodes.SendErrorJSON(w, r, http.StatusNotFound, err)
			return
		}
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}
