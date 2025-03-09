package error_handler

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql" // Import for MySQL error handling
	"github.com/lib/pq"              // Import for PostgreSQL error handling
	"github.com/mattn/go-sqlite3"    // Import for SQLite error handling
	"gorm.io/gorm"
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

// mapDBError maps database errors to custom error types
func mapDBError(err error) *Error {
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, sql.ErrNoRows) {
		e := NotFoundError("Database record not found")
		return &e
	}

	if IsUniqueConstraintViolation(err) {
		e := ConflictError("Duplicate entry, unique constraint violated")
		return &e
	}

	// Instead of returning InternalServerError, return nil
	return nil
}

// IsUniqueConstraintViolation checks for unique constraint violations
func IsUniqueConstraintViolation(err error) bool {
	// PostgreSQL unique constraint violation (pq error code "23505")
	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
		return true
	}

	// MySQL unique constraint violation (Error code 1062)
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		return true
	}

	// SQLite unique constraint violation
	if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
		return true
	}

	return false
}
