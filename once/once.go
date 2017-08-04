package once

func GenerateToken(url string) (token string, err error) {
	token = url
	return
}

func IsTokenUsed(token string) (isUsed bool, err error) {
	isUsed = true
	return
}

func SetTokenAsUsed(token string) (err error) {
	return
}