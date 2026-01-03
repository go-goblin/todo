# URL Shortener Service

HTTP сервис для сокращения URL с REST API и редиректами.

## Описание

Сервис предоставляет два основных эндпоинта:
1. **Создание короткой ссылки** - принимает оригинальный URL, возвращает короткий код
2. **Редирект по короткой ссылке** - перенаправляет пользователя на оригинальный URL

## API Эндпоинты

### 1. Создание короткой ссылки

**Запрос:**
```http
POST /create
Content-Type: application/json

{
    "url": "https://example.com/very/long/url/path"
}
```

**Ответ:**

```json
{
    "data": {
        "short_code": "abc123"
    }
}
```

### 1. Создание короткой ссылки

**Запрос:**
```http
GET /abc123
```

**Ответ:**
- HTTP 302 Found с редиректом на оригинальный URL
- Или HTTP 404 Not Found если код не существует


### Использование
#### Создание короткой ссылки

```bash
# Linux/Mac
curl -X POST http://localhost:8080/create \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com"}'

# Windows PowerShell
curl.exe -X POST http://localhost:8080/create `
  -H "Content-Type: application/json" `
  -d '{"url": "https://google.com"}'
```

#### Переход по короткой ссылке

```bash
curl -L http://localhost:8080/a1b2c3
```

### Структура проекта
```text
urlshortener/
├── cmd/server/          # Точка входа сервера
├── internal/
│   ├── handler/         # HTTP обработчики
│   ├── service/         # Бизнес-логика
│   ├── models/          # Структуры данных
│   └── dependencies/    # DI контейнер
└── go.mod
```

### Запуск
```bash
# Установка зависимостей
go mod tidy

# Запуск сервера
go run cmd/server/main.go

# Сервер запустится на http://localhost:8080
```

### Маршрутизация
- GET /{code} - редирект на оригинальный URL
- POST /create - создание новой короткой ссылки

### Особенности реализации
- Генерация 6-символьных кодов (буквы a-z, цифры 0-9)
- Хранение данных в памяти (in-memory)
- Автоматическое добавление схемы https:// если отсутствует
- Валидация входных данных
- RESTful JSON API

### Пример рабочего процесса
1. Пользователь отправляет длинный URL на /create
2. Сервис генерирует короткий код (например abc123)
3. Сервис сохраняет связку abc123 → оригинальный URL
4. Пользователь переходит по http://your-domain.com/abc123
5. Сервис находит оригинальный URL и делает 302 редирект
6. Пользователь попадает на целевую страницу