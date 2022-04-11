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

	refreshToID = "https://securetoken.googleapis.com/v1/token?key=%s"
)

var apiKey = os.Getenv("API_KEY")

func SignInWithCustomToken2(Customtoken string) (entity.Token, error) {
	var token entity.Token

	request, err := json.Marshal(map[string]interface{}{
		"token":             Customtoken,
		"returnSecureToken": true,
	})
	if err != nil {
		return entity.Token{}, err
	}

	response, err := postRequest(fmt.Sprintf(verifyCustomToken, apiKey), request)
	if err != nil {
		return entity.Token{}, err
	}

	if err := json.Unmarshal(response, &token); err != nil {
		return entity.Token{}, err
	}

	return token, nil
}

func postRequest(url string, request []byte) ([]byte, error) {
	response, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

func SignInWithCustomToken(Customtoken string) (entity.Token, error) {
	httpposturl := fmt.Sprintf(verifyCustomToken, apiKey)

	requestJSON, err := json.Marshal(map[string]interface{}{
		"token":             Customtoken,
		"returnSecureToken": true,
	})
	if err != nil {
		return entity.Token{}, err
	}

	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(requestJSON))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var token entity.Token
	if err := json.Unmarshal(body, &token); err != nil {
		return entity.Token{}, err
	}
	return token, nil
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

//https://firebase.google.com/docs/reference/rest/auth?hl=en#section-verify-custom-token
