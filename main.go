package main

import (
	"firebase-authentication/routes"
	"firebase-authentication/service"
)

func main() {
	// Setup Gin-Gonic
	service.StartFirebase()
	routes.GinSetup()
}
