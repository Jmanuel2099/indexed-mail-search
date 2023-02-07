package customerror

import (
	"net/http"

	"github.com/go-chi/render"
)

type CustomError struct {
	Status    int    `json:"status"`
	ErrorText string `json:"error,omitempty"`
}

func NewCustomError(statusCode int, errMessage string) *CustomError {
	customErr := &CustomError{
		Status:    statusCode,
		ErrorText: errMessage,
	}
	return customErr
}

func (cr *CustomError) ErrorResponseHandling(w http.ResponseWriter, r *http.Request) {
	render.Status(r, cr.Status)
	render.JSON(w, r, cr)
}
