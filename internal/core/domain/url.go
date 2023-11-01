package domain

import (
	"math/rand"
	"time"
)

type Url struct {
	Id   int
	Url  string
	Code string
}

func NewUrl(url string) Url {
	return Url{
		Url:  url,
		Code: getCode(),
	}
}

func getCode() string {
	rand.Seed(time.Now().UnixNano())

	allowedChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	randomChars := make([]byte, 6)

	for i := 0; i < 6; i++ {
		randomChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}

	return string(randomChars)
}

func (url Url) GetShortUrl() string {
	return "httsp://my-short-url.com/" + url.Code
}
