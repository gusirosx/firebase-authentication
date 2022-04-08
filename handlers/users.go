package handlers

import (
	"firebase-authentication/models"
	"firebase-authentication/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user struct {
	Uid string `json:"uid"`
}

func Token(ctx *gin.Context) {

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Unmarshalling Data"})
	}
	tokenString, err := service.Client.CustomToken(ctx, user.Uid)
	if err != nil {
		fmt.Println(err.Error())
	}
	idToken, err := models.SignInWithCustomToken(tokenString)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"idToken": idToken})
}
