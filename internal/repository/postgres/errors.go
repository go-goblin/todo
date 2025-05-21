package postgres

import "errors"

var (
	ErrSelect       = errors.New("ошибка выборки из базы данных")
	ErrCreateUser   = errors.New("ошибка создания пользователя")
	ErrCreateTask   = errors.New("ошибка создания задачи")
	ErrUpdateTask   = errors.New("ошибка обновления задачи")
	ErrDeleteTask   = errors.New("ошибка удаления задачи")
	ErrTaskNotFound = errors.New("задача не найдена")
)
