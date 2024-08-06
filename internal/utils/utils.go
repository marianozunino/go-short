package utils

import (
	"crypto"
	"fmt"
	"math/rand"
	"net/http"
)

func IsValidURL(url string) bool {
	_, err := http.Get(url)
	return err == nil
}

func GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func Md5(s string) string {
	hash := crypto.MD5.New()
	output := hash.Sum([]byte(s))
	return fmt.Sprintf("%x", output)
}
