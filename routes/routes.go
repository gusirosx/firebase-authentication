package routes

import (
	"firebase-authentication/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port string = ":8000"

func GinSetup() {
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Initialize API routes
	InitializeRoutes(router)

	// Start serving the application
	if err := router.Run(port); err != nil {
		log.Fatalln("could not run the application: ", err.Error())
	} else {
		log.Fatalf("Server listening on port" + string(port))
	}
}

func InitializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})
	// Handle Get requests for ID token's
	router.POST("/token", handlers.GetToken)
	router.POST("/refreshtoken", handlers.RefreshIDtoken)

	// Handle the no route case
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})
}
