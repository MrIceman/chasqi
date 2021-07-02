package processor

import (
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
