package agent

import "testing"

func extractHeadersShouldRaiseAnErrorIfParametersWereNotSeen(t *testing.T) {
	subject := Agent{}

	subject.inputValues = map[string]string{
		"name": "email",
		"blah": "eh",
	}
}
