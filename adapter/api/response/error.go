package response

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

var (
	ErrParameterInvalid = errors.New("parameter invalid")

	ErrInvalidInput = errors.New("invalid input")
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessage(message []string, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     message,
	}
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}
