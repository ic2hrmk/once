package once

import (
	"strings"
	"encoding/hex"
	"crypto/rand"
	"github.com/garyburd/redigo/redis"
	"path"
)

const (
	HTTP_PREFIX = "http://"
	HTTPS_PREFIX = "https://"
)

const (
	SALT_LENGTH = 4
)

func GenerateShortLink(url string) (shortLink string, err error) {
	if !strings.Contains(url, HTTP_PREFIX) && !strings.Contains(url, HTTPS_PREFIX) {
		url = HTTP_PREFIX + url
	}

	token := generateRandomString()

	conn := redisPool.Get()
	_, err = conn.Do("SET", token, url)
	if err != nil {
		return
	}

	defer conn.Close()

	shortLink = path.Join(domain, token)

	return
}

func GetShortLinkValue(token string) (url string, err error) {
	conn := redisPool.Get()
	url, err = redis.String(conn.Do("GET", token))

	defer conn.Close()

	// If token not found - it's mean that it's used
	if err != nil {
		return
	}

	return
}

func IsShortLinkUsed(token string) (isUsed bool) {
	conn := redisPool.Get()
	value, err := redis.String(conn.Do("GET", token))

	defer conn.Close()

	// If token not found - it's mean that it's used
	if err != nil || value == "" {
		isUsed = true
	}

	return
}

func SetShortLinkAsUsed(token string) (err error) {
	conn := redisPool.Get()
	_, err = conn.Do("DEL", token)
	if err != nil {
		return
	}
	defer conn.Close()

	return
}

func generateRandomString() string {
	data := make([]byte, SALT_LENGTH)
	rand.Read(data)
	return hex.EncodeToString(data[:])
}