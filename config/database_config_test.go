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
		"DB_USER":    "postgres",
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

	dbConfig := getDatabaseConfig()
	assert.Equal(t, dbConfigVars["DB_USER"], dbConfig.user)
	assert.Equal(t, dbConfigVars["DB_NAME"], dbConfig.dbname)
	assert.Equal(t, dbConfigVars["DB_HOST"], dbConfig.host)
	assert.Equal(t, dbConfigVars["DB_PORT"], dbConfig.port)
	assert.Equal(t, dbConfigVars["ENABLE_SSL"], dbConfig.ssl)

	testDbConfigEnd()
}

func TestDatabaseConnString(t *testing.T) {
	testDbConfigInit()
	assert.Equal(t, "user=postgres dbname=userland host=localhost port=5432 sslmode=disable", GetDatabaseConnectionString())
	testDbConfigEnd()
}
