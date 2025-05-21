package service

type SignIn struct {
	Username string
	Password string
}

type CreateTask struct {
	Title       string
	Description string
	Status      string
	UserID      int
}

type UpdateTask struct {
	ID          int
	Title       string
	Description string
	Status      string
	UserID      int
}
