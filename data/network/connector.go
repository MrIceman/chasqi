package network

import (
	"bytes"
	"chasqi/data/auth"
	"encoding/json"
	"net/http"
)

type NetworkConnector struct {
	auth *auth.AuthApi
}

func (nc *NetworkConnector) Init(api *auth.AuthApi) {
	nc.auth = api
}

func (nc *NetworkConnector) GetAuthenticated(url string) (*http.Response, error) {
	token := nc.auth.SessionToken

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 401 {
		println("Refresh required")
		nc.auth.Refresh()
		println("Refreshed")
		return nc.GetAuthenticated(url)
	}

	return res, err
}

func (nc *NetworkConnector) PostAuthenticated(url string, body interface{}) (*http.Response, error) {
	token := nc.auth.SessionToken

	client := &http.Client{}
	reqBody, err := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 401 {
		println("Refresh required")
		nc.auth.Refresh()
		println("Refreshed")
		return nc.PostAuthenticated(url, body)
	}

	return res, err
}

func (nc *NetworkConnector) PutAuthenticated(url string, body interface{}) (*http.Response, error) {
	token := nc.auth.SessionToken

	client := &http.Client{}
	reqBody, err := json.Marshal(body)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 401 {
		println("Refresh required")
		nc.auth.Refresh()
		println("Refreshed")
		return nc.PostAuthenticated(url, body)
	}

	return res, err
}
