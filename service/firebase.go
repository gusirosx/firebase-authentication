package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// Setup Firebase
var Client *auth.Client

//Firebase admin SDK initialization
func GetClientFirebase() (*auth.Client, error) {
	// Initialize default app
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Println("Error initializing the app" + string(err.Error()))
	}
	// Access auth service from the default app
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Println("Error getting Auth client" + string(err.Error()))
	}
	return client, err
}
