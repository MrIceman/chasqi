package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestShouldReturnARandomStringWith10Chars(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	result := GenerateRandomString(10)
	println(result)
	if len(result) != 10 {
		t.Error("String doesn't have 10 chars")
	}
}
