package models

type BaseResponse struct {
	Error *BaseError `json:"error,omitempty"`
}

type BaseError struct {
	Message string `json:"message,omitempty"`
}

type ShortUrlResultData struct {
	ShortCode string `json:"short_code"`
}

type ShortUrlResultResponse struct {
	BaseResponse
	Data *ShortUrlResultData `data:"error,omitempty"`
}
