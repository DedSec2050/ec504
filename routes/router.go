package routes

import (
	"eclab/handlers" // Update import path to 'eclab/handlers'

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the Gin routes
func SetupRoutes(r *gin.Engine) {
	// Define routes for each sample text using the handler
	r.GET("/sample/:id", handlers.GetSampleText)
}
