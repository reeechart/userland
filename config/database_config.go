package config

import (
	"fmt"
	"os"
)

type databaseConfig struct {
	user   string
	dbname string
	host   string
	port   string
	ssl    string
}

const (
	DB_NAME    = "userland"
	DB_HOST    = "localhost"
	DB_PORT    = "5432"
	ENABLE_SSL = "disable"
)

func GetDatabaseConnectionString() string {
	dbConfig := getDatabaseConfig()
	return fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=%s", dbConfig.user, dbConfig.dbname, dbConfig.host, dbConfig.port, dbConfig.ssl)
}

func getDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		user:   os.Getenv("DB_USER"),
		dbname: os.Getenv("DB_NAME"),
		host:   os.Getenv("DB_HOST"),
		port:   os.Getenv("DB_PORT"),
		ssl:    os.Getenv("ENABLE_SSL"),
	}
}
