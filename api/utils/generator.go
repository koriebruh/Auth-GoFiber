package utils

import (
	"math/rand/v2"
)

func GeneratorRandString(n int) string {
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	letters := make([]rune, n)
	for _, i := range letters {
		letters[i] = charsets[rand.IntN(len(charsets))]
	}

	return string(letters)

}
