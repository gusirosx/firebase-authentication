package handlers

import (
	"bytes"
	"encoding/json"
	"firebase-authentication/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Token(ctx *gin.Context) {
	var user struct {
		Uid string `json:"uid"`
	}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Unmarshalling Data"})
	}
	tokenString, err := service.Client.CustomToken(ctx, user.Uid)
	if err != nil {
		fmt.Println(err.Error())
	}
	idToken, err := signInWithCustomToken(tokenString)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"idToken": idToken})
}

const (
	verifyCustomTokenURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s"
)

var (
	apiKey = os.Getenv("API_KEY")
)

func signInWithCustomToken(token string) (string, error) {
	req, err := json.Marshal(map[string]interface{}{
		"token":             token,
		"returnSecureToken": true,
	})
	if err != nil {
		return "", err
	}

	resp, err := postRequest(fmt.Sprintf(verifyCustomTokenURL, apiKey), req)
	if err != nil {
		return "", err
	}
	var respBody struct {
		IDToken string `json:"idToken"`
	}
	if err := json.Unmarshal(resp, &respBody); err != nil {
		return "", err
	}
	return respBody.IDToken, err
}

func postRequest(url string, req []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
