package common

import (
	"bytes"
	"net/http"
)

// ResponseRecorder для перехвата ответа в мидлварах
type ResponseRecorder struct {
	http.ResponseWriter
	Status int
	Size   int64
	Body   *bytes.Buffer
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		Body:           &bytes.Buffer{},
	}
}

// WriteHeader Переопределение для сохранения реального статуса ответа
func (r *ResponseRecorder) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

// Write Переопределение для сохранения размера ответа
func (r *ResponseRecorder) Write(b []byte) (int, error) {
	if r.Body != nil {
		r.Body.Write(b)
	}
	size, err := r.ResponseWriter.Write(b)
	r.Size += int64(size)
	return size, err
}
