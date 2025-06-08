# ToDo REST API

Простое REST API для управления задачами с авторизацией через JWT. Поддерживает регистрацию пользователей, логин и CRUD-операции над задачами.

---

## 🚀 Возможности

- Регистрация пользователя
- Авторизация и генерация JWT
- Создание, обновление, удаление и получение задач
- Ограничение доступа к задачам других пользователей
- Swagger-документация

---

## 🧱 Сервисная структура

### UserService
Обрабатывает регистрацию и вход пользователей:
- `SignIn` — регистрация с хешированием пароля и выдачей JWT
- `Login` — проверка пароля, выдача токена
- Использует `bcrypt` и `TokenManager`

### TaskService
Работает с задачами:
- `GetAllTasks` / `GetTaskByID` — чтение
- `CreateTask` / `UpdateTask` / `DeleteTask` — изменение задач
- Валидация прав доступа (владелец == пользователь)

---

## 📦 Запуск проекта

```bash
cd devops/local
docker compose up
```

- После запуска API будет доступно по адресу: http://localhost:8080
- Swagger-документация: http://localhost:8080/docs

## 📂 Структура

- `cmd/server/main.go` — точка входа
- `internal/service/` — бизнес-логика (UserService, TaskService)
- `internal/repository/` — слой доступа к данным (PostgreSQL)
- `internal/models/` — DTO и структуры ответов/запросов
- `internal/config/` — конфиг + env переменные
- `pkg/auth/ `— JWT генерация и валидация
- `docs/` — сгенерированная Swagger-документация (swag init)