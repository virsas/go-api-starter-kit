package helpers

import (
	"math/rand"
	"time"
)

func RandomString(size int, prefix string) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, size)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	randomString := prefix + string(str)

	return randomString
}
