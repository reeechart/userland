package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRespondSuccess(t *testing.T) {
	expectedBodyContent := map[string]bool{"success": true}
	expectedBody, err := json.Marshal(expectedBodyContent)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondSuccess(res)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, string(expectedBody), res.Body.String(), "Respond success should return success true JSON")
}

func TestRespondSuccessWithBody(t *testing.T) {
	sampleObject := map[string]interface{}{
		"time":    time.Now(),
		"weather": "sunny",
	}
	expectedBody, err := json.Marshal(sampleObject)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondSuccessWithBody(res, sampleObject)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, string(expectedBody), res.Body.String())
}
