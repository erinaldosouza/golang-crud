package app_errors

import "net/http"

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AppError struct {
	Message string   `json:"message"`
	Err   string   `json:"error"`
	Code    int      `json:"code"`
	Cause   []Causes `json:"causes"`
}

func (r *AppError) Error() string {
	return r.Message
}

func NewAppError(message, error string, code int, cause []Causes) *AppError {
    return &AppError {
        Message: message,
        Err:     error,
        Code:    code,
        Cause:   cause,
    }
}

func NewNotFoundError(message string, cause []Causes) *AppError {
    return NewAppError(message, "Resource not found",  http.StatusNotFound, cause)
}

func NewBadRequestError(message string, cause []Causes) *AppError {
    return NewAppError(message, "Bad Request", http.StatusBadRequest, cause)
}

func NewValidationError(message string, cause []Causes) *AppError {
    return NewAppError(message, "Validation Error", http.StatusUnprocessableEntity, cause)
}

func NewNotAuthorizedError(message string, cause []Causes) *AppError {
    return NewAppError(message, "Not Authorized", http.StatusUnauthorized, cause)
}

func NewForbiddenError(message string, cause []Causes) *AppError {
    return NewAppError(message, "Forbidden", http.StatusForbidden, cause)
}