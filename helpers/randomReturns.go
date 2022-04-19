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

func RandomInt64(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())

	value := rand.Int63n(max-min+1) + min
	return value
}

func RandomPubString(length int, size int) string {
	var finalString string

	for i := 1; i <= size; i++ {
		finalString = finalString + RandomString(length, "")
		if i < size {
			finalString = finalString + "-"
		}
	}

	return finalString
}
