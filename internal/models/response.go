package models

type BaseResponse struct {
	Error *BaseError `json:"error,omitempty"`
}

type BaseError struct {
	Message string `json:"message,omitempty"`
}

type OperationResultResponse struct {
	BaseResponse
	Data *OperationResultData `data:"error,omitempty"`
}

type OperationResultData struct {
	Success bool `json:"success"`
}

type AuthResponse struct {
	BaseResponse
	Data *AuthData `data:"error,omitempty"`
}

type AuthData struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

type TaskListResponse struct {
	BaseResponse
	Data *[]TaskDTO `data:"error,omitempty"`
}

type TaskResponse struct {
	BaseResponse
	Data *TaskDTO `data:"error,omitempty"`
}

type CreateTaskResponse struct {
	BaseResponse
	Data *CreateTaskData `json:"data,omitempty"`
}

type CreateTaskData struct {
	TaskID int `json:"task_id"`
}
