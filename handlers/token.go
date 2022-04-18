package handlers

import (
	"firebase-authentication/entity"
	"firebase-authentication/models"
	"firebase-authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetToken(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Unmarshalling Data"})
		return
	}

	tokenString, err := service.Client.CustomToken(ctx, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := models.SignInWithCustomToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": token})
}

func RefreshIDtoken(ctx *gin.Context) {

	RefreshToken := ctx.GetHeader("refresh_token")

	if RefreshToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Unmarshalling Data"})
		return
	}

	token, err := models.RefreshIDtoken(RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": token})
}

func SendPasswordResetEmail(ctx *gin.Context) {

	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Unmarshalling Data"})
		return
	}

	if err := models.SendPasswordResetEmail(user.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "ok"})

}
