package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	ulanderrors "userland/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	SAMPLE_ERROR_CODE             = 1000
	SAMPLE_ERROR_MESSAGE          = "sample error"
	APPLICATION_JSON_CONTENT_TYPE = "application/json"
)

var (
	sampleError = ulanderrors.UserlandError{Code: SAMPLE_ERROR_CODE, Message: SAMPLE_ERROR_MESSAGE}
)

func TestRespondSuccess(t *testing.T) {
	expectedBodyContent := map[string]bool{"success": true}
	expectedBody, err := json.Marshal(expectedBodyContent)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondSuccess(res)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, APPLICATION_JSON_CONTENT_TYPE, res.Header().Get("Content-Type"))
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
	assert.Equal(t, APPLICATION_JSON_CONTENT_TYPE, res.Header().Get("Content-Type"))
	assert.Equal(t, string(expectedBody), res.Body.String())
}

func TestRespondBadRequest(t *testing.T) {
	expectedBody, err := json.Marshal(sampleError)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondBadRequest(res, sampleError)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, APPLICATION_JSON_CONTENT_TYPE, res.Header().Get("Content-Type"))
	assert.Equal(t, string(expectedBody), res.Body.String())
}

func TestRespondUnauthorized(t *testing.T) {
	expectedBody, err := json.Marshal(sampleError)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondUnauthorized(res, sampleError)
	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.Equal(t, APPLICATION_JSON_CONTENT_TYPE, res.Header().Get("Content-Type"))
	assert.Equal(t, string(expectedBody), res.Body.String())
}

func TestRespondInternalError(t *testing.T) {
	expectedBody, err := json.Marshal(sampleError)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondInternalError(res, sampleError)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, APPLICATION_JSON_CONTENT_TYPE, res.Header().Get("Content-Type"))
	assert.Equal(t, string(expectedBody), res.Body.String())
}
