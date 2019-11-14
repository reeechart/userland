package config

const (
	JWT_KEY = "userland_jwt_key"
)

func GetJWTKey() string {
	return JWT_KEY
}
