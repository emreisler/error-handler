package error_handler

import (
	"net/http"
)

type Error struct {
	message    string
	httpStatus int
}

func (e Error) Error() string {
	return e.message
}

func (e Error) HttpStatus() int {
	return e.httpStatus
}

func New(message string, httpStatus int) Error {
	return Error{message, httpStatus}
}

// Predefined Web Service Errors

func InternalServerError(msg string) Error {
	return New(msg, http.StatusInternalServerError)
}

func BadRequestError(msg string) Error {
	return New(msg, http.StatusBadRequest)
}

func UnauthorizedError(msg string) Error {
	return New(msg, http.StatusUnauthorized)
}

func ForbiddenError(msg string) Error {
	return New(msg, http.StatusForbidden)
}

func NotFoundError(msg string) Error {
	return New(msg, http.StatusNotFound)
}

func ConflictError(msg string) Error {
	return New(msg, http.StatusConflict)
}

func UnprocessableEntityError(msg string) Error {
	return New(msg, http.StatusUnprocessableEntity)
}

func TooManyRequestsError(msg string) Error {
	return New(msg, http.StatusTooManyRequests)
}

func ServiceUnavailableError(msg string) Error {
	return New(msg, http.StatusServiceUnavailable)
}
