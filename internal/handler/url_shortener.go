package handler

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"strings"
	"url-stortener/internal/errcodes"
	"url-stortener/internal/httpx"
	"url-stortener/internal/service"
)

func (h *Handler) RedirectToUrl(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	if code == "" {
		errcodes.SendErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("short code is required"))
		return
	}

	originalURL, found := h.deps.UrlShortenerService.RedirectToUrl(code)
	if !found {
		http.NotFound(w, r)
		return
	}

	fullURL := ensureScheme(originalURL)

	http.Redirect(w, r, fullURL, http.StatusFound)
}

func (h *Handler) PostMakeShortUrl(w http.ResponseWriter, r *http.Request) {
	var request MakeShortUrlRequest
	if err := httpx.DecodeAndValidateBody(w, r, &request); err != nil {
		return
	}
	input := &service.CreateShortUrl{
		Url: request.Url,
	}
	response, err := h.deps.UrlShortenerService.CreateShortUrl(input)
	if err != nil {
		errcodes.SendErrorJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, response)
}

// ensureScheme добавляет https:// если URL не имеет схемы
func ensureScheme(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url // Уже есть схема
	}
	// Добавляем https:// по умолчанию
	return "https://" + url
}
