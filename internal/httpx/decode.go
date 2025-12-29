package httpx

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"url-stortener/internal/models"
)

// DecodeAndValidateBody декодирует тело JSON запроса в целевую структуру и валидирует её.
// Метод автоматически отправляет HTTP 400 Bad Request с JSON ошибкой при неудаче.
//
// Пример использования с MakeShortUrlRequest:
//
//	    // 1. Объявляем переменную для хранения декодированных данных
//	    var req models.MakeShortUrlRequest
//
//	    // 2. Вызываем метод, передавая writer, request и указатель на структуру
//	    if err := DecodeAndValidateBody(w, r, &req); err != nil {
//	        // Метод уже отправил ответ с ошибкой, можно просто вернуться
//	        // или добавить логирование:
//	        log.Printf("Failed to process request: %v", err)
//	        return
//	    }
//	}
func DecodeAndValidateBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.BaseResponse{
				Error: &models.BaseError{
					Message: "failed to decode request",
				}})
			return err
		}
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, models.BaseResponse{
			Error: &models.BaseError{
				Message: "failed to decode request",
			},
		})
		return err
	}

	validate := validator.New()
	if err := validate.Struct(dst); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, models.BaseResponse{
			Error: &models.BaseError{
				Message: "validation failed",
			},
		})
		return err
	}
	return nil
}
