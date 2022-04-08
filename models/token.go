package models

import (
	"bytes"
	"encoding/json"
	"firebase-authentication/entity"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const verifyCustomTokenURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s"

var apiKey = os.Getenv("API_KEY")

func SignInWithCustomToken(Customtoken string) (entity.Token, error) {
	var token entity.Token

	request, err := json.Marshal(map[string]interface{}{
		"token":             Customtoken,
		"returnSecureToken": true,
	})
	if err != nil {
		return entity.Token{}, err
	}

	response, err := postRequest(fmt.Sprintf(verifyCustomTokenURL, apiKey), request)
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
