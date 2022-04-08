package routes

import (
	"firebase-authentication/handlers"
	"firebase-authentication/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func GinSetup() {
	const port string = ":8080"
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	if err := router.Run(port); err != nil {
		err := fmt.Errorf("could not run the application: %v", err)
		log.Fatalf(err.Error())
	} else {
		log.Fatalf("Server listening on port" + string(port))
	}
}

func initializeRoutes() {
	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	//router.Use(middleware.SetUserStatus())
	// Handle the index route
	router.GET("/", func(c *gin.Context) {
		// Contact the server and print out its response.
		fmt.Fprintln(c.Writer, "Up and running...")
	})

	// Setup Firebase
	var err error
	service.Client, err = service.GetClientFirebase()
	if err != nil {
		err := fmt.Errorf("error getting the auth client: %v", err)
		log.Fatalf(err.Error())
	}
	// Handle token
	router.GET("/token", handlers.Token)
}
