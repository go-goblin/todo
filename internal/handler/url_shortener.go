package handler

import (
	"github.com/go-chi/render"
	"net/http"
	"todo/internal/errcodes"
	"todo/internal/httpx"
	"todo/internal/service"
)

func (h *Handler) PostMakeShortUrl(w http.ResponseWriter, r *http.Request) {
	var request MakeShortUrlRequest
	if err := httpx.DecodeAndValidateBody(w, r, &request); err != nil {
		return
	}
	input := &service.CreateTask{
		Title: request.Url,
	}
	response, err := h.deps.TaskService.CreateTask(input)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}
