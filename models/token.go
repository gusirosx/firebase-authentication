package models

import (
	"bytes"
	"encoding/json"
	"firebase-authentication/entity"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	verifyCustomTokenURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s"

	verifyCustomToken = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s"

	passwordResetEmail = "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=%s"

	refreshToID = "https://securetoken.googleapis.com/v1/token?key=%s"
)

var apiKey = os.Getenv("API_KEY")

func SignInWithCustomToken(Customtoken string) (entity.Token, error) {
	request, err := json.Marshal(map[string]interface{}{
		"token":             Customtoken,
		"returnSecureToken": true,
	})
	if err != nil {
		return entity.Token{}, err
	}

	response, err := postRequest(fmt.Sprintf(verifyCustomToken, apiKey), "application/json", request)
	if err != nil {
		return entity.Token{}, err
	}

	var token entity.Token
	if err := json.Unmarshal(response, &token); err != nil {
		return entity.Token{}, err
	}
	return token, nil
}

func postRequest(url string, contentType string, request []byte) ([]byte, error) {
	response, err := http.Post(url, contentType, bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

//Exchange a refresh token for an ID token
func RefreshIDtoken(refreshToken string) (entity.RefreshToken, error) {
	endpoint := fmt.Sprintf(refreshToID, apiKey)
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	client := &http.Client{}
	response, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	response.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(response)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	var token entity.RefreshToken
	if err := json.Unmarshal(body, &token); err != nil {
		return entity.RefreshToken{}, err
	}
	return token, nil

}

func SendPasswordResetEmail(email string) error {
	request, err := json.Marshal(map[string]interface{}{
		"requestType": "PASSWORD_RESET",
		"email":       email,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = postRequest(fmt.Sprintf(passwordResetEmail, apiKey), "application/json", request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// var token entity.Token
	// if err := json.Unmarshal(response, &token); err != nil {
	// 	return entity.Token{}, err
	// }
	return nil
}

// =============================================================================
// https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=[API_KEY]
// curl 'https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=[API_KEY]' \
// -H 'Content-Type: application/json' \
// --data-binary '{"requestType":"PASSWORD_RESET","email":"[user@example.com]"}'

// {
//  "email": "[user@example.com]"
// }

// =============================================================================
// curl 'https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=[API_KEY]' \
// -H 'Content-Type: application/json' \
// --data-binary '{"token":"[CUSTOM_TOKEN]","returnSecureToken":true}'

//https://firebase.google.com/docs/reference/rest/auth?hl=en#section-verify-custom-token

// request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(requestJSON))
// request.Header.Set("Content-Type", "application/json; charset=UTF-8")

// client := &http.Client{}
// response, error := client.Do(request)
// if error != nil {
// 	panic(error)
// }
// defer response.Body.Close()

// body, _ := ioutil.ReadAll(response.Body)
