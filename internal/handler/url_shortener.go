package handler

import (
	"net/http"
	"strings"
)

func (h *Handler) RedirectToUrl(w http.ResponseWriter, r *http.Request) {
	// TODO: Извлечь короткий код из URL пути
	// TODO: Проверить что код не пустой
	// TODO: Если код пустой, вернуть ошибку 400 Bad Request через errcodes.SendErrorJSON
	// TODO: Получить оригинальный URL по коду из UrlShortenerService.RedirectToUrl
	// TODO: Если URL не найден (found == false), вернуть 404 Not Found
	// TODO: Обеспечить наличие схемы (http/https) в URL с помощью ensureScheme
	// TODO: Выполнить HTTP редирект на оригинальный URL с кодом 302 Found
}

func (h *Handler) PostMakeShortUrl(w http.ResponseWriter, r *http.Request) {
	// TODO: Декодировать и валидировать тело JSON запроса в структуру MakeShortUrlRequest
	// TODO: Если ошибка декодирования/валидации, вернуть ошибку (метод уже отправляет ответ)
	// TODO: Создать структуру service.CreateShortUrl из данных запроса
	// TODO: Вызвать метод UrlShortenerService.CreateShortUrl для создания короткой ссылки
	// TODO: Вернуть успешный JSON ответ с данными короткой ссылки через render.JSON
}

// ensureScheme добавляет https:// если URL не имеет схемы
func ensureScheme(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url // Уже есть схема
	}
	// Добавляем https:// по умолчанию
	return "https://" + url
}
