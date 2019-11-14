package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dbConfigVars map[string]string
)

func testDbConfigInit() {
	dbConfigVars = map[string]string{
		"DB_NAME":    "userland",
		"DB_HOST":    "localhost",
		"DB_PORT":    "5432",
		"ENABLE_SSL": "disable",
	}

	for key, val := range dbConfigVars {
		os.Setenv(key, val)
	}
}

func testDbConfigEnd() {
	for key := range dbConfigVars {
		os.Unsetenv(key)
	}
}

func TestDatabaseConfig(t *testing.T) {
	testDbConfigInit()

	dbConfig, err := getDatabaseConfig()
	assert.Nil(t, err)
	assert.Equal(t, dbConfig.dbname, dbConfigVars["DB_NAME"])
	assert.Equal(t, dbConfig.host, dbConfigVars["DB_HOST"])
	assert.Equal(t, dbConfig.port, dbConfigVars["DB_PORT"])
	assert.Equal(t, dbConfig.ssl, dbConfigVars["ENABLE_SSL"])

	testDbConfigEnd()
}

func TestDatabaseConnString(t *testing.T) {
	testDbConfigInit()
	assert.Equal(t, GetDatabaseConnectionString(), "dbname=userland host=localhost port=5432 sslmode=disable")
	testDbConfigEnd()
}
