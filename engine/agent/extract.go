package agent

import (
	"fmt"
	"strings"
)

func (a *Agent) extractHeaders(route *Route) (map[string]string, error) {
	resultMap := make(map[string]string)
	for k, v := range route.headers {
		values := strings.Split(fmt.Sprint(v), "::")
		if len(values) > 1 {
			isVariable := values[0] == "variable"
			if isVariable {
				resultMap[k] = fmt.Sprint(a.inputValues[values[1]])
			} else {
				// output from previous response
				exposedValue := a.exposedResponses[values[0]][values[1]]
				resultMap[k] = exposedValue
				if k == "Authorization" {
					resultMap[k] = "Bearer " + exposedValue
				}
			}
		} else {
			resultMap[k] = fmt.Sprint(v)
		}
	}

	return resultMap, nil
}

func (a *Agent) extractBody(route *Route) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})

	for k, v := range route.body {
		values := strings.Split(fmt.Sprint(v), "::")
		if len(values) > 1 {
			// TODO handle exceptions
			isVariable := values[0] == "variables"
			if isVariable {
				resultMap[k] = a.inputValues[values[1]]
			} else {
				// output from previous response
				exposedValue := a.exposedResponses[values[0]][values[1]]
				// todo handle error if exposedValue is noll
				resultMap[k] = exposedValue
			}
		} else {
			resultMap[k] = fmt.Sprint(v)
		}
	}

	return resultMap, nil
}
