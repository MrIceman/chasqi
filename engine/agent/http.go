package agent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (a *Agent) makeGetOrDeleteRequest(
	method string,
	route *Route,
	headers map[string]string) *http.Request {
	if method != GET && method != DELETE {
		panic(errors.New("Method must be GET or DELETE and not " + method))
	}
	req, err := http.NewRequest("GET",
		a.host+route.path,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// set up headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req
}

func (a *Agent) makePostOrPutRequest(
	method string,
	route *Route,
	headers map[string]string,
	requestBody map[string]interface{}) *http.Request {
	if method != POST && method != PUT {
		panic(errors.New("Method must be PUT or POST and not " + method))
	}
	body, err := json.Marshal(requestBody)
	a.sendDebugMessage("response body: " + fmt.Sprintln(string(body)))
	if err != nil {
		panic(err)
	}

	// create the request body
	req, err := http.NewRequest(method,
		a.host+route.path,
		bytes.NewBuffer(body),
	)
	if err != nil {
		panic(err)
	}

	// set up headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req
}
