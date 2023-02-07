package customerror

import (
	"net/http"

	"github.com/go-chi/render"
)

// CustomError it is a struc that will be used when you want to handle
// an error in a personalized way.
type CustomError struct {
	Status    int    `json:"status"`
	ErrorText string `json:"error,omitempty"`
}

// NewCustomError works as the conntrucutor of the CustomError struc
func NewCustomError(statusCode int, errMessage string) *CustomError {
	customErr := &CustomError{
		Status:    statusCode,
		ErrorText: errMessage,
	}
	return customErr
}

// ErrorResponseHandling function that is implemented when responding
// to an http request with an error
func (cr *CustomError) ErrorResponseHandling(w http.ResponseWriter, r *http.Request) {
	render.Status(r, cr.Status)
	render.JSON(w, r, cr)
}
