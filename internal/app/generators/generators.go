package generators

import (
	"crypto/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SecureRandomString(length int) string {
	b := make([]byte, length)

	randomData := make([]byte, length)
	if _, err := rand.Read(randomData); err != nil {
		panic(err) // В продакшене надо возвращать ошибку
	}

	for i := range b {
		b[i] = charset[int(randomData[i])%len(charset)]
	}

	return string(b)
}
