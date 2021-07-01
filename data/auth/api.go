package auth

import (
	"bytes"
	"chasqi/data/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type AuthApi struct {
	SessionToken     *config.SessionToken
	Host             string
	ObtainTokenPath  string
	RefreshTokenPath string
}

func (a *AuthApi) LogIn(credentials config.Credentials) {
	fmt.Println("Performing Login")
	var url strings.Builder
	body, err := json.Marshal(credentials)
	if err != nil {
		panic("Could not marshal credentials")
	}
	url.WriteString(a.Host)
	url.WriteString(a.ObtainTokenPath)
	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic("Failed to perform request, " + err.Error())
	}
	fmt.Println(strconv.Itoa(resp.StatusCode))

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	token := config.SessionToken{}

	if err != nil {
		panic("Failed to decode byte body, " + err.Error())
	}

	json.Unmarshal(bytes, &token)
	fmt.Println(token.AccessToken)
	fmt.Println(token.RefreshToken)

	a.SessionToken = &token
}

func (a *AuthApi) Refresh() {
	fmt.Println("Performing refresh")
	var url strings.Builder
	body, err := json.Marshal(a.SessionToken)
	if err != nil {
		panic("Could not marshal credentials")
	}
	url.WriteString(a.Host)
	url.WriteString(a.RefreshTokenPath)
	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic("Failed to perform request, " + err.Error())
	}
	fmt.Println(strconv.Itoa(resp.StatusCode))

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	println(string(bytes))
	token := config.SessionToken{}

	if err != nil {
		panic("Failed to decode byte body, " + err.Error())
	}

	json.Unmarshal(bytes, &token)
	fmt.Println(token.AccessToken)
	fmt.Println(token.RefreshToken)

	a.SessionToken = &token
}
