package main

import (
	"github.com/emreisler/error-handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Apply global error handler middleware
	r.Use(error_handler.ErrorHandlerMiddleware())

	// Define routes
	r.GET("/example", ExampleHandler)

	// Start the server
	r.Run(":8080")
}

// Example handler that returns different errors
func ExampleHandler(c *gin.Context) {
	// Simulate an error
	err := error_handler.BadRequestError("Invalid input provided")

	// Instead of returning response here, we just call c.Error()
	c.Error(err)
}
