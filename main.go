package main

import (
	"firebase-authentication/routes"
	"firebase-authentication/service"
)

func main() {
	// Initialize Firebase Auth client.
	service.StartFirebase()
	// Setup Gin-Gonic server.
	routes.GinSetup()
}
