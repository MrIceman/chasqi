package agent

import (
	"chasqi/engine/analytics"
	"chasqi/rules"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Agent struct {
	identifier              int
	alreadyVisitedRoutes    map[string]string
	exposedResponses        map[string]map[string]string
	inputValues             map[string]interface{}
	rootRoute               Route
	sleepTimeInSeconds      int
	host                    string
	debugChannel            chan string
	logChannel              chan analytics.LogEntry
	client                  http.Client
	lastExecutedRequestBody map[string]interface{}
}

func New(sleepTimeInSeconds int) Agent {
	a := Agent{sleepTimeInSeconds: sleepTimeInSeconds}

	return a
}

func (a *Agent) Identifier() int {
	return a.identifier
}

func (a *Agent) Init(tree rules.NavigationTree,
	debugChannel chan string,
	logChannel chan analytics.LogEntry,
	identifier int,
) {
	a.rootRoute = *nodesToRoute(tree.Nodes)
	a.logChannel = logChannel
	a.host = tree.Host
	a.debugChannel = debugChannel

	inputMap := make(map[string]interface{})
	// Pass a copy of the configured variables to the agent
	for _, item := range tree.Variables {
		if item.Type == "str" {
			if strings.Contains(item.Value, "random::") {
				inputMap[item.Name] = ReplaceRandomsInString(item.Value)
			} else {
				inputMap[item.Name] = item.Value
			}
		} else if item.Type == "int" {
			intValue, err := strconv.Atoi(item.Value)
			if err != nil {
				panic(err)
			}
			inputMap[item.Name] = intValue
		}
	}
	a.identifier = identifier
	a.inputValues = inputMap
	a.exposedResponses = make(map[string]map[string]string)
	for _, n := range tree.Nodes {
		a.exposedResponses[n.Name] = make(map[string]string)
		for _, response := range n.ExposesResponse {
			// set an empty response map to the agent with
			// the values
			// create empty default map
			// create empty default value for the required response
			// it will be replaced later when actual the value is obtained
			a.exposedResponses[n.Name][response] = ""
		}
	}
}

func (a *Agent) sendDebugMessage(mesage string) {
	a.debugChannel <- "{" + strconv.Itoa(a.identifier) + "} " + mesage
}

// The agent starts his route
func (a *Agent) Start() {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(3)
	time.Sleep(time.Duration(randInt+1) * time.Second)
	a.sendDebugMessage("starting")
	currentRoute := &a.rootRoute
	for currentRoute != nil {
		resp := a.callRoute(currentRoute)
		// See if we have to cache any values
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		a.processResponse(bytes, currentRoute)

		currentRoute = currentRoute.next

		_ = resp.Body.Close()
		time.Sleep(2 * time.Second)
	}

}

func (a *Agent) processResponse(data []byte, r *Route) {
	if !r.returnsArray {
		a.processSingleValueResponse(data, r)
	} else {
		a.processMultiValueResponse(data, r)
	}
}

// Parses a response that returns a single value
func (a *Agent) processSingleValueResponse(data []byte, r *Route) {
	var resultBody map[string]interface{}
	err := json.Unmarshal(data, &resultBody)
	if err != nil {
		panic(err)
	}

	for route, exposeValues := range a.exposedResponses {
		// check if the current route is within our expose map
		if route == r.name {
			for key, _ := range exposeValues {
				value := resultBody[key]
				if value == nil {
					panic("No value found for " + key + " when calling " + r.name)
				}
				a.exposedResponses[r.name][key] = fmt.Sprint(value)
			}
		}
	}
}

// Parses a response that returns an array
func (a *Agent) processMultiValueResponse(data []byte, r *Route) {
	var resultBody []map[string]interface{}
	err := json.Unmarshal(data, &resultBody)
	if err != nil {
		a.sendDebugMessage("Could not parse response: " + err.Error())
	}
}

func (a *Agent) callRoute(route *Route) *http.Response {
	method := route.method
	var req *http.Request
	var headerMap, err = a.extractHeaders(route)
	if err != nil {
		panic(err)
	}
	if method == POST || method == PUT {
		var bodyMap, err = a.extractBody(route)
		if err != nil {
			panic(err)
		}
		a.lastExecutedRequestBody = bodyMap
		req = a.makePostOrPutRequest(
			method,
			route,
			headerMap,
			bodyMap)
	}
	if method == GET || method == DELETE {
		req = a.makeGetOrDeleteRequest(
			method,
			route,
			headerMap)
	}

	startTime := time.Now()
	response, err := a.client.Do(req)
	endTime := time.Now()
	if err != nil {
		panic(err)
	}
	durationTimeInMs := int(endTime.Sub(startTime).Milliseconds())
	a.sendDebugMessage(strconv.Itoa(response.StatusCode) +
		" - " + route.method + " " + route.name + "  / " +
		strconv.Itoa(durationTimeInMs) + " ms" + " - ")

	a.logChannel <- analytics.LogEntry{
		response.StatusCode,
		route.name,
		durationTimeInMs,
		nil,
		route.method,
	}
	return response
}
