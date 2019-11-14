package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	JWT_KEY = "userland_jwt_key"
)

func GetJWTKey() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return os.Getenv("JWT_KEY")
}
