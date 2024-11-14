package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewUnexpected(msg string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewBadRequestError(msg string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}
