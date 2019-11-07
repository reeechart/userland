package config

import (
	"fmt"
)

const (
	DB_NAME    = "userland"
	DB_HOST    = "localhost"
	DB_PORT    = "5432"
	ENABLE_SSL = "disable"
)

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("dbname=%s host=%s port=%s sslmode=%s", DB_NAME, DB_HOST, DB_PORT, ENABLE_SSL)
}
