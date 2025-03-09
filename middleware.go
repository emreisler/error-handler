package error_handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// ErrorHandlerMiddleware handles all errors centrally
func ErrorHandlerMiddleware() gin.HandlerFunc {
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

			// If it's an unknown error, return 500
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Internal Server Error",
				"status": http.StatusInternalServerError,
			})
		}
	}
}
