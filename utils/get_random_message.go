package utils

import (
	"fmt"
	"math/rand"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomMessage(n int) string {
	message := ""
	for length := 0; length < n; length++ {
		text := make([]byte, n)
		for i := range text {
			text[i] = charset[rand.Intn(len(charset))]
		}

		message = fmt.Sprintf("%s %s", message, text)
	}

	return message
}
