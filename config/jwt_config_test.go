package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_JWT_KEY = "test_jwt_key"
)

func testJWTConfigInit() {
	os.Setenv("JWT_KEY", TEST_JWT_KEY)
}

func testJWTConfigEnd() {
	os.Unsetenv("JWT_KEY")
}

func TestJWTKey(t *testing.T) {
	testJWTConfigInit()
	assert.Equal(t, TEST_JWT_KEY, GetJWTKey())
	testJWTConfigEnd()
}
