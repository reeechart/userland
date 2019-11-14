package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	router *mux.Router
)

func testPingInit() {
	router = mux.NewRouter()
	router.HandleFunc("/ping", Ping).Methods(http.MethodGet)
}

func TestPing(t *testing.T) {
	testPingInit()

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.Nil(t, err)

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}
