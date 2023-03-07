package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var (
	Client *auth.Client
)

// StartFirebase initializes the Firebase Admin SDK client.
func StartFirebase() {
	var err error
	Client, err = GetClientFirebase()
	if err != nil {
		log.Fatalln("error getting the auth client: ", err.Error())
	}
}

// GetClientFirebase returns an authenticated Firebase Auth client.
func GetClientFirebase() (*auth.Client, error) {
	// Initialize the default Firebase app.
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Get the Auth client from the default Firebase app.
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
