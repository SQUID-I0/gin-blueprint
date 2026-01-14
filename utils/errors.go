package utils

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrBadRequest        = errors.New("bad request")
	ErrInternalServer    = errors.New("internal server error")
	ErrValidation        = errors.New("validation error")
	ErrDuplicateEntry    = errors.New("duplicate entry")
	ErrDatabaseOperation = errors.New("database operation failed")
)

type AppError struct {
	Err        error
	Message    string
	Code       string
	StatusCode int
	Details    map[string]interface{}
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Err.Error()
}

func NewAppError(err error, message, code string, statusCode int) *AppError {
	return &AppError{
		Err:        err,
		Message:    message,
		Code:       code,
		StatusCode: statusCode,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Err:        ErrNotFound,
		Message:    message,
		Code:       "NOT_FOUND",
		StatusCode: http.StatusNotFound,
	}
}

func NewValidationError(message string, details map[string]interface{}) *AppError {
	return &AppError{
		Err:        ErrValidation,
		Message:    message,
		Code:       "VALIDATION_ERROR",
		StatusCode: http.StatusBadRequest,
		Details:    details,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Err:        ErrUnauthorized,
		Message:    message,
		Code:       "UNAUTHORIZED",
		StatusCode: http.StatusUnauthorized,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Err:        ErrInternalServer,
		Message:    message,
		Code:       "INTERNAL_ERROR",
		StatusCode: http.StatusInternalServerError,
	}
}

func NewDuplicateError(message string) *AppError {
	return &AppError{
		Err:        ErrDuplicateEntry,
		Message:    message,
		Code:       "DUPLICATE_ENTRY",
		StatusCode: http.StatusConflict,
	}
}
