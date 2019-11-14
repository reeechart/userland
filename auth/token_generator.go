package auth

import (
	"math/rand"
	"time"
	"userland/config"

	"github.com/dgrijalva/jwt-go"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

const (
	TOKEN_CHARS  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TOKEN_LENGTH = 32

	HOURS_IN_DAY = 24
)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func generateToken() string {
	token := make([]byte, TOKEN_LENGTH)
	for i := range token {
		token[i] = TOKEN_CHARS[seededRand.Intn(len(TOKEN_CHARS))]
	}
	return string(token)
}

func generateJWT(user User, expirationTime time.Time) (string, error) {
	claims := Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetJWTKey()))

	return tokenString, err
}
