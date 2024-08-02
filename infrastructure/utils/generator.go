package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(value int) string {
	rand.Seed(time.Now().UnixNano())
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letters := make([]rune, value)
	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(letters)
}