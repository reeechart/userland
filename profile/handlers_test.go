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

	validEmailReq   ChangeEmailRequest
	invalidEmailReq ChangeEmailRequest

	validChangePassReq          ChangePasswordRequest
	invalidPassChangePassReq    ChangePasswordRequest
	unmatchingPassChangePassReq ChangePasswordRequest

	deleteAccReq DeleteAccountRequest
)

func testProfileHandlerInit(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = NewMockprofileRepositoryInterface(ctrl)

	handler = ProfileHandler{ProfileRepo: mockRepo}

	router = mux.NewRouter()
	router.HandleFunc("/api/me", handler.GetProfile).Methods(http.MethodGet)
	router.HandleFunc("/api/me", handler.UpdateProfile).Methods(http.MethodPut)
	router.HandleFunc("/api/me/email", handler.GetEmail).Methods(http.MethodGet)
	router.HandleFunc("/api/me/email", handler.ChangeEmailAddress).Methods(http.MethodPut)
	router.HandleFunc("/api/me/password", handler.ChangePassword).Methods(http.MethodPost)
	router.HandleFunc("/api/me/delete", handler.DeleteAccount).Methods(http.MethodPost)
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
	req, err := http.NewRequest(http.MethodGet, "/api/me/email", nil)
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
	assert.Equal(t, "{\"email\":\""+user.Email+"\"}", res.Body.String())
}

func TestChangeEmailAddress(t *testing.T) {
	testProfileHandlerInit(t)
	initSuiteAndRepoForChangeEmail()

	testChangeUserEmail(t, &authenticatedUser, validEmailReq, http.StatusOK)
	testChangeUserEmail(t, &authenticatedUser, invalidEmailReq, http.StatusBadRequest)

	testProfileHandlerEnd()
}

func initSuiteAndRepoForChangeEmail() {
	validEmailReq = ChangeEmailRequest{
		NewEmail: "changedemail@example.com",
	}

	invalidEmailReq = ChangeEmailRequest{
		NewEmail: "invalidemailaddress",
	}

	mockRepo.EXPECT().changeUserEmail(&authenticatedUser, validEmailReq.NewEmail).Return(nil)
}

func testChangeUserEmail(t *testing.T, user *auth.User, emailReq ChangeEmailRequest, expectedStatusCode int) {
	newEmailData, err := json.Marshal(emailReq)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPut, "/api/me/email", bytes.NewReader(newEmailData))
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}

func TestChangePassword(t *testing.T) {
	testProfileHandlerInit(t)
	initSuiteAndRepoForChangePassword()

	testChangeUserPassword(t, &authenticatedUser, validChangePassReq, http.StatusOK)
	testChangeUserPassword(t, &authenticatedUser, invalidPassChangePassReq, http.StatusBadRequest)
	testChangeUserPassword(t, &authenticatedUser, unmatchingPassChangePassReq, http.StatusBadRequest)

	testProfileHandlerEnd()
}

func initSuiteAndRepoForChangePassword() {
	validChangePassReq = ChangePasswordRequest{
		PasswordCurrent: "password",
		Password:        "passwordnew",
		PasswordConfirm: "passwordnew",
	}

	invalidPassChangePassReq = ChangePasswordRequest{
		PasswordCurrent: "password",
		Password:        "pass",
		PasswordConfirm: "pass",
	}

	unmatchingPassChangePassReq = ChangePasswordRequest{
		PasswordCurrent: "password",
		Password:        "newpassword",
		PasswordConfirm: "othernewpassword",
	}

	mockRepo.EXPECT().changeUserPassword(&authenticatedUser, validChangePassReq.PasswordCurrent, validChangePassReq.Password).Return(nil)
}

func testChangeUserPassword(t *testing.T, user *auth.User, changePassReq ChangePasswordRequest, expectedStatusCode int) {
	changePassData, err := json.Marshal(changePassReq)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPost, "/api/me/password", bytes.NewReader(changePassData))
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}

func TestDeleteAccount(t *testing.T) {
	testProfileHandlerInit(t)
	mockRepo.EXPECT().deleteUser(&authenticatedUser, authenticatedUser.Password).Return(nil)

	testDeleteUserAccount(t, &authenticatedUser, http.StatusOK)

	testProfileHandlerEnd()
}

func testDeleteUserAccount(t *testing.T, user *auth.User, expectedStatusCode int) {
	deleteAccData, err := json.Marshal(deleteAccReq)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPost, "/api/me/delete", bytes.NewReader(deleteAccData))
	req = setRequestUserContext(req, user)
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}
