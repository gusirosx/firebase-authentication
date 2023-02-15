// Package routes provides the API routes for the application.
package routes

import (
	"firebase-authentication/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = ":8000"

// GinSetup initializes and runs the API server using Gin.
func GinSetup() {
	// Create a new router instance
	router := gin.Default()

	// Initialize API routes
	initializeRoutes(router)

	// Start serving the application
	if err := router.Run(port); err != nil {
		log.Fatalf("could not run the application: %v", err)
	} else {
		log.Printf("Server listening on port %v", port)
	}
}

// initializeRoutes sets up the routes for the API server.
func initializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})

	// Handle token-related routes
	router.POST("/token", handlers.GetToken)
	router.POST("/refreshtoken", handlers.RefreshIDtoken)

	// Handle password reset email route
	router.POST("/passwordResetEmail", handlers.SendPasswordResetEmail)

	// Handle the no route case
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})
}
