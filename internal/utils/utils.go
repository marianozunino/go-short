package utils

import (
	"crypto"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
)

type Result struct {
	IsValid bool
	Message string
}

func IsValidURL(target string) Result {
	if target == "" {
		return Result{IsValid: false, Message: "URL cannot be empty"}
	}

	u, err := url.ParseRequestURI(target)
	if err != nil {
		return Result{IsValid: false, Message: "Invalid URL provided"}
	}

	s, err := http.Get(u.String())

	if err != nil {
		return Result{IsValid: false, Message: "We couldn't reach the server"}
	}

	if s.StatusCode >= 400 {
		return Result{IsValid: false, Message: "Not found"}
	}

	return Result{IsValid: true, Message: target}
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
