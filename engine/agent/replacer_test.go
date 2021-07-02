package agent

import (
	"strings"
	"testing"
)

func TestShouldReplaceSimlpeRandomString(t *testing.T) {
	str := "random::{3}"

	result := ReplaceRandomsInString(str)
	println(result)
	if len(result) != 3 {
		t.Error("Generated string doesn't have 3 chars")
	}
}

func TestShouldReplaceSimpleRandomStringWithPrefix(t *testing.T) {
	str := "1234onerandom::{3}"

	result := ReplaceRandomsInString(str)
	println(result)
	if len(result) != 10 {
		t.Error("Generated string doesn't have 3 chars")
	}
}

func TestShouldReplaceStringWithMultipleRandoms(t *testing.T) {
	str := "random::{3}@random::{150}.com"

	result := ReplaceRandomsInString(str)
	println(result)
	if len(result) != 3+1+150+4 {
		t.Error("Generated string doesn't have 3 chars")
	}
}

func TestShouldContainInfixString(t *testing.T) {
	str := "random::{1200}infixrandom::{150}"

	result := ReplaceRandomsInString(str)

	if !strings.Contains(result, "infix") {
		t.Error("Infix \"infix\" is not contained")
	}
	if len(result) != 1355 {
		t.Error("Result should contain 1355 chars")
	}
}
