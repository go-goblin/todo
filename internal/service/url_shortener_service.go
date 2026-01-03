package service

import (
	"time"
	"url-stortener/internal/models"
)

type UrlShortenerService struct {
	usedCodes  map[string]bool   // Для проверки уникальности кода
	urlStorage map[string]string // Для хранения: код -> оригинальный URL
}

func NewUrlShortenerService() UrlShortenerService {
	return UrlShortenerService{
		usedCodes:  make(map[string]bool),
		urlStorage: make(map[string]string),
	}
}

func (s *UrlShortenerService) RedirectToUrl(code string) (string, bool) {
	// TODO: Найти оригинальный URL по короткому коду в хранилище urlStorage
	// TODO: Если URL найден (exists == true), вернуть URL и true
	// TODO: Если URL не найден, вернуть пустую строку и false
	return "", false
}

func (s *UrlShortenerService) CreateShortUrl(input *CreateShortUrl) *models.ShortUrlResultResponse {
	// TODO: Сгенерировать короткий код для переданного URL с помощью generateCode
	// TODO: Сохранить связку код → URL в хранилище urlStorage
	// TODO: Создать и вернуть структуру ShortUrlResultResponse с сгенерированным кодом
	// TODO: Вернуть nil ошибку так как операция всегда успешна в текущей реализации
	return nil
}

// generateCode создает уникальный 6-символьный код для короткой ссылки
func (s *UrlShortenerService) generateCode(url string) string {
	// Используем только строчные буквы и цифры (36 символов)
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"

	// Генерируем код пока не найдем уникальный
	for {
		code := ""
		for i := 0; i < 6; i++ {
			// Используем время как источник случайности
			charIndex := int(time.Now().UnixNano()) % len(chars)
			code += string(chars[charIndex])
			time.Sleep(1 * time.Nanosecond) // Меняем время для следующего символа
		}

		// Проверяем уникальность кода
		if !s.usedCodes[code] {
			s.usedCodes[code] = true
			return code
		}

		// Если код уже существует, пробуем снова
	}
}
