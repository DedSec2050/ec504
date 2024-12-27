package main

import (
	"eclab/routes" // Update import path to use 'eclab/routes'

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Set up the routes
	routes.SetupRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
