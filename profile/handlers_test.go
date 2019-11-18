package profile

import (
	"bytes"
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

	validProfileUpdate           UserProfile
	invalidFullnameProfileUpdate UserProfile
	invalidLocationProfileUpdate UserProfile
	invalidBioProfileUpdate      UserProfile
	invalidWebProfileUpdate      UserProfile
)

func testProfileHandlerInit(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = NewMockprofileRepositoryInterface(ctrl)

	handler = ProfileHandler{ProfileRepo: mockRepo}

	router = mux.NewRouter()
	router.HandleFunc("/api/me", handler.GetProfile).Methods(http.MethodGet)
	router.HandleFunc("/api/me", handler.UpdateProfile).Methods(http.MethodPut)
	router.HandleFunc("/api/me/email", handler.GetEmail).Methods(http.MethodGet)
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

func TestUpdateProfile(t *testing.T) {
	testProfileHandlerInit(t)
	initSuiteAndRepoForUpdateProfile()

	testUpdateUserProfile(t, &authenticatedUser, validProfileUpdate, http.StatusOK)
	testUpdateUserProfile(t, &authenticatedUser, invalidFullnameProfileUpdate, http.StatusBadRequest)
	testUpdateUserProfile(t, &authenticatedUser, invalidLocationProfileUpdate, http.StatusBadRequest)
	testUpdateUserProfile(t, &authenticatedUser, invalidBioProfileUpdate, http.StatusBadRequest)
	testUpdateUserProfile(t, &authenticatedUser, invalidWebProfileUpdate, http.StatusBadRequest)

	testProfileHandlerEnd()
}

func initSuiteAndRepoForUpdateProfile() {
	validProfileUpdate = UserProfile{
		Fullname: "updateduser",
		Location: "Jakarta, Indonesia",
		Bio:      "my new bio",
		Web:      "https://example.com/newme",
	}

	invalidFullnameProfileUpdate = UserProfile{
		Fullname: "un",
		Location: "Jakarta, Indonesia",
		Bio:      "my new bio",
		Web:      "https://example.com/newme",
	}

	invalidLocationProfileUpdate = UserProfile{
		Fullname: "updateduser",
		Location: STR_LEN_MORE_THAN_128,
		Bio:      "my new bio",
		Web:      "https://example.com/newme",
	}

	invalidBioProfileUpdate = UserProfile{
		Fullname: "updateduser",
		Location: "Jakarta, Indonesia",
		Bio:      STR_LEN_MORE_THAN_128 + STR_LEN_MORE_THAN_128,
		Web:      "https://example.com/newme",
	}

	invalidWebProfileUpdate = UserProfile{
		Fullname: "updateduser",
		Location: "Jakarta, Indonesia",
		Bio:      "my new bio",
		Web:      "whatiwanttofill",
	}

	mockRepo.EXPECT().updateUserProfile(&authenticatedUser, validProfileUpdate).Return(nil)
}

func testUpdateUserProfile(t *testing.T, user *auth.User, profileUpdate UserProfile, expectedStatusCode int) {
	profileUpdateData, err := json.Marshal(profileUpdate)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPut, "/api/me", bytes.NewReader(profileUpdateData))
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}

func TestGetEmail(t *testing.T) {
	testProfileHandlerInit(t)

	testGetUserEmail(t, &authenticatedUser, http.StatusOK)

	testProfileHandlerEnd()
}

func testGetUserEmail(t *testing.T, user *auth.User, expectedStatusCode int) {
	_, err := json.Marshal(user)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodGet, "/api/me/email", nil)
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}
