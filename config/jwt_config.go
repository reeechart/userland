package config

import (
	"os"
)

const (
	JWT_KEY = "userland_jwt_key"
)

func GetJWTKey() string {
	return os.Getenv("JWT_KEY")
}
