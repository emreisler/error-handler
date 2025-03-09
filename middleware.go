package error_handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GinMiddleware handles all errors centrally
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// Recover from panic if one occurs
			if r := recover(); r != nil {
				slog.ErrorContext(c, "[PANIC RECOVERED] %v", r)

				// Return a 500 response for panics
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Internal Server Error",
					"status": http.StatusInternalServerError,
				})
				c.Abort()
			}
		}()

		c.Next() // Process request

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			var serviceErr Error
			if errors.As(err, &serviceErr) {
				// If the error matches our custom error type, return structured response
				c.JSON(serviceErr.HttpStatus(), gin.H{
					"error":  serviceErr.Error(),
					"status": serviceErr.HttpStatus(),
				})
				return
			}

			if dbErr := mapDBError(err.Err); dbErr != nil {
				c.JSON(dbErr.HttpStatus(), gin.H{
					"error":  dbErr.Error(),
					"status": dbErr.HttpStatus(),
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Internal Server Error",
				"status": http.StatusInternalServerError,
			})
		}
	}
}

// Helper struct to capture errors in net/http
type errorResponseWriter struct {
	http.ResponseWriter
	err error
}

func (w *errorResponseWriter) WriteHeader(statusCode int) {
	if statusCode >= 400 {
		w.err = errors.New(http.StatusText(statusCode))
	}
	w.ResponseWriter.WriteHeader(statusCode)
}

// Net/HTTP Middleware
func NetHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC RECOVERED] %v", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Capture response writer for handling errors
		errorWriter := &errorResponseWriter{ResponseWriter: w}
		next.ServeHTTP(errorWriter, r)

		if errorWriter.err != nil {
			// Check for service error
			var serviceErr Error
			if errors.As(errorWriter.err, &serviceErr) {
				http.Error(w, serviceErr.Error(), serviceErr.HttpStatus())
				return
			}

			// Check for database error
			if dbErr := mapDBError(errorWriter.err); dbErr != nil {
				http.Error(w, dbErr.Error(), dbErr.HttpStatus())
				return
			}

			// Default to 500 internal server error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

// Chi Middleware for centralized error handling
func ChiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC RECOVERED] %v", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Capture response writer for handling errors
		errorWriter := &errorResponseWriter{ResponseWriter: w}
		next.ServeHTTP(errorWriter, r)

		if errorWriter.err != nil {
			// Check for service error
			var serviceErr Error
			if errors.As(errorWriter.err, &serviceErr) {
				http.Error(w, serviceErr.Error(), serviceErr.HttpStatus())
				return
			}

			// Check for database error
			if dbErr := mapDBError(errorWriter.err); dbErr != nil {
				http.Error(w, dbErr.Error(), dbErr.HttpStatus())
				return
			}

			// Default to 500 internal server error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

// FiberMiddleware for centralized error handling
func FiberMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC RECOVERED] %v", r)
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error", "status": fiber.StatusInternalServerError})
			}
		}()

		err := c.Next()

		if err != nil {
			var serviceErr Error
			if errors.As(err, &serviceErr) {
				return c.Status(serviceErr.HttpStatus()).JSON(fiber.Map{"error": serviceErr.Error(), "status": serviceErr.HttpStatus()})
			}

			if dbErr := mapDBError(err); dbErr != nil {
				return c.Status(dbErr.HttpStatus()).JSON(fiber.Map{"error": dbErr.Error(), "status": dbErr.HttpStatus()})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error", "status": fiber.StatusInternalServerError})
		}

		return nil
	}
}
