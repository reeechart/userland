package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	res                  *httptest.ResponseRecorder
	userRegistrationData []byte

	handler AuthHandler
	router  *mux.Router

	ctrl     *gomock.Controller
	mockRepo *MockuserRepositoryInterface

	validNewUser          userRegistration
	invalidNewUser        userRegistration
	incompleteNewUser     userRegistration
	unmatchingPassNewUser userRegistration
	validNewUserFailQuery userRegistration
)

func testAuthHandlerInit(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = NewMockuserRepositoryInterface(ctrl)

	handler = AuthHandler{UserRepo: mockRepo}

	router = mux.NewRouter()
	router.HandleFunc("/auth/register", handler.Register).Methods(http.MethodPost)
}

func testAuthHandlerEnd() {
	ctrl.Finish()
}

func initRepoForRegistration() {
	validNewUser = userRegistration{
		Fullname:        "user",
		Email:           "user@example.com",
		Password:        "password",
		PasswordConfirm: "password",
	}

	invalidNewUser = userRegistration{
		Fullname:        "user",
		Email:           "email_invalid@examplecom",
		Password:        "password",
		PasswordConfirm: "password",
	}

	incompleteNewUser = userRegistration{
		Fullname: "user",
		Email:    "user@example.com",
	}

	unmatchingPassNewUser = userRegistration{
		Fullname:        "user",
		Email:           "user@example.com",
		Password:        "password",
		PasswordConfirm: "differentpassword",
	}

	validNewUserFailQuery = userRegistration{
		Fullname:        "user invalid query",
		Email:           "user@example.com",
		Password:        "password",
		PasswordConfirm: "password",
	}

	gomock.InOrder(
		mockRepo.EXPECT().createNewUser(validNewUser).Return(nil),
		// mockRepo.EXPECT().createNewUser(validNewUserFailQuery).Return(errors.New("")),
	)
}

func TestRegister(t *testing.T) {
	testAuthHandlerInit(t)
	initRepoForRegistration()

	userRegistrationData, err = json.Marshal(validNewUser)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(userRegistrationData))
	require.Nil(t, err)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	userRegistrationData, err = json.Marshal(invalidNewUser)
	require.Nil(t, err)
	req, err = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(userRegistrationData))
	require.Nil(t, err)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)

	userRegistrationData, err = json.Marshal(incompleteNewUser)
	require.Nil(t, err)
	req, err = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(userRegistrationData))
	require.Nil(t, err)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)

	testAuthHandlerEnd()
}
