package random

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	// setting a random seed
	rand.Seed(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	letterBuffer := make([]byte, length)

	for i := range letterBuffer {
		randomInt := rand.Intn(len(letterBytes))
		letterBuffer[i] = letterBytes[randomInt]
	}

	return string(letterBuffer)
}
