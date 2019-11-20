package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	middleware AuthMiddleware

	authenticatedUser = User{
		Id:    1,
		Email: "user@example.com",
	}

	nextHandler = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	validReq         *http.Request
	cookielessReq    *http.Request
	temperedTokenReq *http.Request
	expiredTokenReq  *http.Request
)

func testAuthMiddlewareInit(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = NewMockuserRepositoryInterface(ctrl)

	middleware = AuthMiddleware{UserRepo: mockRepo}

	router = mux.NewRouter()
	router.HandleFunc("/with/auth", middleware.WithVerifyJWT(nextHandler)).Methods(http.MethodGet)
}

func testAuthMiddlewareEnd() {
	ctrl.Finish()
}

func TestWithVerifyJWT(t *testing.T) {
	testAuthMiddlewareInit(t)
	initSuiteAndRepoForVerifyJWT(t)

	testJWTVerificationRequest(t, validReq, http.StatusOK)
	testJWTVerificationRequest(t, cookielessReq, http.StatusUnauthorized)
	testJWTVerificationRequest(t, temperedTokenReq, http.StatusUnauthorized)
	testJWTVerificationRequest(t, expiredTokenReq, http.StatusUnauthorized)

	testAuthMiddlewareEnd()
}

func initSuiteAndRepoForVerifyJWT(t *testing.T) {
	expirationTime := time.Now().Add(HOURS_IN_DAY * time.Hour)
	jwtToken, err := generateJWT(authenticatedUser, expirationTime)
	require.Nil(t, err)

	validTokenCookie := http.Cookie{
		Name:    "token",
		Value:   jwtToken,
		Expires: expirationTime,
	}
	validReq, _ = http.NewRequest(http.MethodGet, "/with/auth", nil)
	validReq.AddCookie(&validTokenCookie)

	cookielessReq, _ = http.NewRequest(http.MethodGet, "/with/auth", nil)

	temperedToken := jwtToken[:len(jwtToken)-1]
	invalidTokenCookie := http.Cookie{
		Name:    "token",
		Value:   temperedToken,
		Expires: expirationTime,
	}
	temperedTokenReq, _ = http.NewRequest(http.MethodGet, "/with/auth", nil)
	temperedTokenReq.AddCookie(&invalidTokenCookie)

	pastTime := time.Now().Add(-1 * HOURS_IN_DAY * time.Hour)
	expiredToken, err := generateJWT(authenticatedUser, pastTime)
	require.Nil(t, err)

	expiredTokenCookie := http.Cookie{
		Name:    "token",
		Value:   expiredToken,
		Expires: pastTime,
	}
	expiredTokenReq, _ = http.NewRequest(http.MethodGet, "/with/auth", nil)
	expiredTokenReq.AddCookie(&expiredTokenCookie)

	mockRepo.EXPECT().getUserById(authenticatedUser.Id).Return(&authenticatedUser, nil)
}

func testJWTVerificationRequest(t *testing.T, req *http.Request, expectedStatusCode int) {
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, expectedStatusCode, res.Code)
}
