package utils

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"fmt"
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
	const length = 6
	bytes := make([]byte, length/2+1)
	if _, err := rand.Read(bytes); err != nil {
		panic(fmt.Sprintf("failed to generate random bytes: %w", err))
	}
	return hex.EncodeToString(bytes)[:length]
}

func Md5(s string) string {
	hash := crypto.MD5.New()
	output := hash.Sum([]byte(s))
	return fmt.Sprintf("%x", output)
}
