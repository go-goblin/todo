package handler

type MakeShortUrlRequest struct {
	Url string `json:"url" validate:"required"`
}
