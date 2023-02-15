package models

import (
	"bytes"
	"encoding/json"
	"firebase-authentication/entity"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//https://firebase.google.com/docs/reference/rest/auth?hl=en#section-verify-custom-token

// Constants
const (
	verifyCustomTokenURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s"
	verifyCustomToken    = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s"
	passwordResetEmail   = "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=%s"
	refreshToID          = "https://securetoken.googleapis.com/v1/token?key=%s"
	defaultContentType   = "application/json"
)

var apiKey = os.Getenv("API_KEY")

// Token represents a Firebase authentication token
type Token struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

// RefreshToken represents a Firebase refresh token
type RefreshToken struct {
	Token         string `json:"id_token"`
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	TokenType     string `json:"token_type"`
	UserID        string `json:"user_id"`
	ProjectID     string `json:"project_id"`
	RefreshType   string `json:"refresh_type"`
	FederatedID   string `json:"federated_id"`
	FederatedIDTS string `json:"federated_id_ts"`
}

// SignInWithCustomToken exchanges a custom token for an authentication token
func SignInWithCustomToken(customToken string) (Token, error) {
	request, err := json.Marshal(map[string]interface{}{
		"token":             customToken,
		"returnSecureToken": true,
	})
	if err != nil {
		return Token{}, err
	}

	response, err := postRequest(fmt.Sprintf(verifyCustomToken, apiKey), defaultContentType, request)
	if err != nil {
		return Token{}, err
	}

	var token Token
	if err := json.Unmarshal(response, &token); err != nil {
		return Token{}, err
	}
	return token, nil
}

// RefreshIDToken exchanges a refresh token for an ID token
func RefreshIDToken(refreshToken string) (RefreshToken, error) {
	endpoint := fmt.Sprintf(refreshToID, apiKey)
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	response, err := postRequest(endpoint, defaultContentType, []byte(data.Encode()))
	if err != nil {
		return RefreshToken{}, err
	}

	var token RefreshToken
	if err := json.Unmarshal(response, &token); err != nil {
		return RefreshToken{}, err
	}
	return token, nil
}

// RefreshIDtoken exchanges a refresh token for an ID token
func RefreshIDtoken(refreshToken string) (entity.RefreshToken, error) {
	endpoint := fmt.Sprintf(refreshToID, apiKey)
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println("Error creating request: ", err)
		return entity.RefreshToken{}, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(request)
	if err != nil {
		log.Println("Error sending request: ", err)
		return entity.RefreshToken{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return entity.RefreshToken{}, fmt.Errorf("unexpected http status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response: ", err)
		return entity.RefreshToken{}, err
	}

	var token entity.RefreshToken
	if err := json.Unmarshal(body, &token); err != nil {
		log.Println("Error unmarshaling response: ", err)
		return entity.RefreshToken{}, err
	}

	return token, nil
}

// SendPasswordResetEmail sends a password reset email to the specified email address
func SendPasswordResetEmail(email string) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"requestType": "PASSWORD_RESET",
		"email":       email,
	})
	if err != nil {
		log.Println("Error creating request: ", err)
		return err
	}

	_, err = postRequest(fmt.Sprintf(passwordResetEmail, apiKey), defaultContentType, requestBody)
	if err != nil {
		log.Println("Error sending request: ", err)
		return err
	}

	return nil
}

func postRequest(url string, contentType string, requestBody []byte) ([]byte, error) {
	response, err := http.Post(url, contentType, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
