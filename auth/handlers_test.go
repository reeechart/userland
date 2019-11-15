package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
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

func TestRegister(t *testing.T) {
	testAuthHandlerInit(t)
	initRepoForRegistration()

	testRegisterUser(t, validNewUser, http.StatusOK)
	testRegisterUser(t, invalidNewUser, http.StatusBadRequest)
	testRegisterUser(t, incompleteNewUser, http.StatusBadRequest)
	testRegisterUser(t, unmatchingPassNewUser, http.StatusBadRequest)
	testRegisterUser(t, validNewUserFailQuery, http.StatusBadRequest)

	testAuthHandlerEnd()
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
		mockRepo.EXPECT().createNewUser(validNewUserFailQuery).Return(errors.New("")),
	)
}

func testRegisterUser(t *testing.T, newUser userRegistration, expectedStatusCode int) {
	userRegistrationData, err := json.Marshal(newUser)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(userRegistrationData))
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}
