package models

import "time"

type TaskDB struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdateAt    time.Time
	UserID      int
}
