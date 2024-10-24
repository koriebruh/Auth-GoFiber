package utils

import (
	"math/rand"
	"time"
)

func GeneratorRandString(n int) string {
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	letters := make([]rune, n)

	// Membuat sumber acak baru
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range letters {
		letters[i] = charsets[r.Intn(len(charsets))]
	}

	return string(letters)
}
