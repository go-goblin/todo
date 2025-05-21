package handler

import (
	"errors"
	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
	"todo/internal/auth"
	"todo/internal/common"
	"todo/internal/errcodes"
	"todo/internal/logger"
	"todo/internal/requestmeta"
)

func (h *Handler) commonWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		commonData := &requestmeta.RequestDTO{
			Method:    r.Method,
			URL:       r.URL.String(),
			StartTime: start,
		}
		ctx := requestmeta.IntoContext(r.Context(), commonData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) handlerLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Получаем общие данные из контекста
		commonData, ok := requestmeta.FromContext(ctx)
		if !ok {
			logger.Get().Error(ctx, "Ошибка получения общих данных из контекста")
			errcodes.SendErrorJSON(w, r, http.StatusBadRequest, errors.New("ошибка получения общих данных из контекста"))
			return
		}
		start := commonData.StartTime

		// Создаем ResponseRecorder для перехвата ответа
		recorder := common.NewResponseRecorder(w)

		// Вызываем следующий обработчик
		next.ServeHTTP(recorder, r)

		// Вычисляем длительность запроса
		duration := time.Since(start).Seconds() // преобразуем в секунды с плавающей точкой

		// Логируем информацию о запросе
		logger.Get().Info(ctx, "Response logger", logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"duration": duration,
			"status":   recorder.Status,
			"size":     recorder.Size,
		})
	})
}

// CorsMiddleware возвращает middleware с настройками CORS.
func (h *Handler) corsMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "X-Requested-With", "Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

func (h *Handler) jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("отсутствует заголовок 'Authorization'"))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, errors.New("ошибка разбора токена"))
			return
		}

		claims := jwt.MapClaims{}
		err := h.deps.TokenManager.ParseToken(token, claims)
		if err != nil {
			errcodes.SendErrorJSON(w, r, http.StatusUnauthorized, err)
			return
		}

		ctx := auth.IntoContext(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
