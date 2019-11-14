package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	SAMPLE_ERROR_CODE = 1000
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

func TestRespondBadRequest(t *testing.T) {
	sampleErr := errors.New("sample error")
	errorResponse := ErrorResponse{Code: SAMPLE_ERROR_CODE, Message: sampleErr.Error()}
	expectedBody, err := json.Marshal(errorResponse)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondBadRequest(res, SAMPLE_ERROR_CODE, sampleErr)
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, string(expectedBody), res.Body.String())
}

func TestRespondUnauthorized(t *testing.T) {
	sampleErr := errors.New("sample unauthorized")
	errorResponse := ErrorResponse{Code: SAMPLE_ERROR_CODE, Message: sampleErr.Error()}
	expectedBody, err := json.Marshal(errorResponse)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondUnauthorized(res, SAMPLE_ERROR_CODE, sampleErr)
	assert.Equal(t, http.StatusUnauthorized, res.Code)
	assert.Equal(t, string(expectedBody), res.Body.String())
}

func TestRespondInternalError(t *testing.T) {
	sampleErr := errors.New("sample internal server error")
	errorResponse := ErrorResponse{Code: SAMPLE_ERROR_CODE, Message: sampleErr.Error()}
	expectedBody, err := json.Marshal(errorResponse)
	require.Nil(t, err)

	res := httptest.NewRecorder()
	RespondInternalError(res, SAMPLE_ERROR_CODE, sampleErr)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, string(expectedBody), res.Body.String())
}
