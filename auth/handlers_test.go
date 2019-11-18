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

	validVerifReq        verificationRequest
	invalidVerifReq      verificationRequest
	invalidTokenVerifReq verificationRequest
	userlessVerifReq     verificationRequest

	loginnableUser   User
	unloginnableUser User
)

const (
	SAMPLE_VALID_VERIFICATION_TOKEN   = "abcdabcdabcdabcdabcdabcdabcd"
	SAMPLE_INVALID_VERIFICATION_TOKEN = "invalidtokeninvalidtokeninvalidt"
)

func testAuthHandlerInit(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = NewMockuserRepositoryInterface(ctrl)

	handler = AuthHandler{UserRepo: mockRepo}

	router = mux.NewRouter()
	router.HandleFunc("/auth/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", handler.Verify).Methods(http.MethodPost)
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

func TestVerify(t *testing.T) {
	testAuthHandlerInit(t)
	initRepoForVerify()

	testVerifyUser(t, validVerifReq, http.StatusOK)
	testVerifyUser(t, invalidVerifReq, http.StatusBadRequest)
	testVerifyUser(t, invalidTokenVerifReq, http.StatusBadRequest)
	testVerifyUser(t, userlessVerifReq, http.StatusBadRequest)

	testAuthHandlerEnd()
}

func initRepoForVerify() {
	validVerifReq = verificationRequest{
		Type:              "email.verify",
		Recipient:         "user@example.com",
		VerificationToken: SAMPLE_VALID_VERIFICATION_TOKEN,
	}

	invalidVerifReq = verificationRequest{
		Type:              "email.verify",
		VerificationToken: SAMPLE_VALID_VERIFICATION_TOKEN,
	}

	invalidTokenVerifReq = verificationRequest{
		Type:              "email.verify",
		Recipient:         "user@example.com",
		VerificationToken: SAMPLE_INVALID_VERIFICATION_TOKEN,
	}

	userlessVerifReq = verificationRequest{
		Type:              "email.verify",
		Recipient:         "notregistered@example.com",
		VerificationToken: SAMPLE_VALID_VERIFICATION_TOKEN,
	}

	gomock.InOrder(
		mockRepo.EXPECT().verifyUser(validVerifReq.Recipient, validVerifReq.VerificationToken).Return(nil),
		mockRepo.EXPECT().verifyUser(invalidTokenVerifReq.Recipient, invalidTokenVerifReq.VerificationToken).Return(errors.New("")),
		mockRepo.EXPECT().verifyUser(userlessVerifReq.Recipient, userlessVerifReq.VerificationToken).Return(errors.New("")),
	)
}

func testVerifyUser(t *testing.T, userVerificationRequest verificationRequest, expectedStatusCode int) {
	verifReq, err := json.Marshal(userVerificationRequest)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPost, "/auth/verify", bytes.NewReader(verifReq))
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}

func TestLogin(t *testing.T) {
	testAuthHandlerInit(t)
	initRepoForLogin()

	testLoginUser(t, loginnableUser, http.StatusOK)
	testLoginUser(t, unloginnableUser, http.StatusBadRequest)

	testAuthHandlerEnd()
}

func initRepoForLogin() {
	loginnableUser = User{
		Email:    "user@example.com",
		Password: "password",
	}

	unloginnableUser = User{
		Email: "user@example.com",
	}

	gomock.InOrder(
		mockRepo.EXPECT().loginUser(loginnableUser.Email, loginnableUser.Password).Return(nil),
		mockRepo.EXPECT().getUserByEmail(loginnableUser.Email).Return(&loginnableUser, nil),
	)
}

func testLoginUser(t *testing.T, loginUser User, expectedStatusCode int) {
	userData, err := json.Marshal(loginUser)
	require.Nil(t, err)
	req, err := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(userData))
	require.Nil(t, err)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}
