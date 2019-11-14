package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type databaseConfig struct {
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
	dbConfig, err := getDatabaseConfig()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("dbname=%s host=%s port=%s sslmode=%s", dbConfig.dbname, dbConfig.host, dbConfig.port, dbConfig.ssl)
}

func getDatabaseConfig() (*databaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &databaseConfig{
		dbname: os.Getenv("DB_NAME"),
		host:   os.Getenv("DB_HOST"),
		port:   os.Getenv("DB_PORT"),
		ssl:    os.Getenv("ENABLE_SSL"),
	}, nil
}
