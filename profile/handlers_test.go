package profile

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"userland/auth"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	handler ProfileHandler
	router  *mux.Router

	ctrl     *gomock.Controller
	mockRepo *MockprofileRepositoryInterface

	authenticatedUser = auth.User{
		Id:             1,
		Fullname:       "userfullname",
		Email:          "user@example.com",
		Location:       sql.NullString{String: "Bandung, Indonesia", Valid: true},
		Bio:            sql.NullString{String: "my bio", Valid: true},
		Web:            sql.NullString{String: "https://example.com", Valid: true},
		ProfilePicture: nil,
		CreatedAt:      time.Now(),
	}
)

func testProfileHandlerInit(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = NewMockprofileRepositoryInterface(ctrl)

	handler = ProfileHandler{ProfileRepo: mockRepo}

	router = mux.NewRouter()
	router.HandleFunc("/api/me", handler.GetProfile).Methods(http.MethodGet)
}

func testProfileHandlerEnd() {
	ctrl.Finish()
}

func setRequestUserContext(req *http.Request, user *auth.User) *http.Request {
	ctx := context.WithValue(req.Context(), "user", user)
	return req.WithContext(ctx)
}

func TestGetProfile(t *testing.T) {
	testProfileHandlerInit(t)
	testGetUserProfile(t, &authenticatedUser, http.StatusOK)
	testProfileHandlerEnd()
}

func testGetUserProfile(t *testing.T, user *auth.User, expectedStatusCode int) {
	_, err := json.Marshal(user)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodGet, "/api/me", nil)
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}
