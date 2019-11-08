package auth

import "math/rand"

const (
	TOKEN_CHARS  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TOKEN_LENGTH = 32
)

func generateToken() string {
	token := make([]byte, TOKEN_LENGTH)
	for i := range token {
		token[i] = TOKEN_CHARS[rand.Intn(len(TOKEN_CHARS))]
	}
	return string(token)
}
