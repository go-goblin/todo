package handler

import (
	"github.com/go-chi/render"
	"net/http"
	"todo/internal/errcodes"
	"todo/internal/httpx"
	"todo/internal/service"
)

// PostSignIn Регистрация
// @Summary Регистрирует пользователя с помощью логина и пароля
// @Tags "auth"
// @Param request body handler.SignInRequest true "Reserve Request Body"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /signup [post]
func (h *Handler) PostSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	var signIn SignInRequest
	if err := httpx.DecodeAndValidateBody(w, r, &signIn); err != nil {
		return
	}
	input := service.SignIn{
		Username: signIn.Username,
		Password: signIn.Password,
	}
	response, err := h.deps.UserService.SignIn(ctx, input)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}

// PostLogin аутентифицирует пользователя по логину и паролю
// @Summary Аутентификация пользователя
// @Tags "auth"
// @Param request body handler.SignInRequest true "Reserve Request Body"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 401 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Router /login [post]
func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	var signIn SignInRequest
	if err := httpx.DecodeAndValidateBody(w, r, &signIn); err != nil {
		return
	}
	input := service.SignIn{
		Username: signIn.Username,
		Password: signIn.Password,
	}
	response, err := h.deps.UserService.Login(ctx, input)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}
