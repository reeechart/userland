package auth

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token := generateToken()
	assert.Equal(t, TOKEN_LENGTH, len(token), "Token should have length of 32 when generated")
	tokenRegex := regexp.MustCompile(`[a-zA-Z0-9]{32}`)
	assert.True(t, tokenRegex.MatchString(token), "Token should only contain lower and uppercased alphabet and numbers")
}
