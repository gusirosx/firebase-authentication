// Package handlers defines HTTP request handlers for the Firebase authentication API
package handlers

import (
	"firebase-authentication/entity"
	"firebase-authentication/models"
	"firebase-authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetToken is a handler function for generating a custom authentication token
// using Firebase client SDK based on the given user ID.
func GetToken(ctx *gin.Context) {
	var user entity.User

	// Bind the JSON data from the request body to the user entity struct
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error unmarshalling data"})
		return
	}

	// Generate a custom authentication token for the user ID using Firebase client SDK
	tokenString, err := service.Client.CustomToken(ctx, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Authenticate the custom token and retrieve the user's ID token
	token, err := models.SignInWithCustomToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the user's ID token in the response body
	ctx.JSON(http.StatusOK, gin.H{"response": token})
}

// RefreshIDtoken exchanges a refresh token for an ID token and returns the ID token in the response
func RefreshIDtoken(ctx *gin.Context) {
	// Get the refresh token from the request header
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		// If the refresh token is missing, return a 400 Bad Request error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token header"})
		return
	}

	// Exchange the refresh token for an ID token
	token, err := models.RefreshIDtoken(refreshToken)
	if err != nil {
		// If an error occurs during the exchange, return a 500 Internal Server Error with the error message
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the ID token in the response body
	ctx.JSON(http.StatusOK, gin.H{"response": token})
}

// SendPasswordResetEmail sends a password reset email to the specified email address
func SendPasswordResetEmail(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error unmarshalling data"})
		return
	}

	if err := models.SendPasswordResetEmail(user.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "Email sent"})
}
