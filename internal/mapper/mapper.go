package mapper

import "todo/internal/models"

func MapTaskDTOFromTaskDb(db models.TaskDB) models.TaskDTO {
	dto := models.TaskDTO{
		ID:          db.ID,
		Title:       db.Title,
		Description: db.Description,
		Status:      db.Status,
		CreatedAt:   db.CreatedAt,
		UpdateAt:    db.UpdateAt,
	}
	return dto
}
