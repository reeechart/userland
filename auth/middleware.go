package auth

import (
	"context"
	"net/http"
	"userland/config"
	ulanderrors "userland/errors"
	"userland/response"

	"github.com/dgrijalva/jwt-go"
)

func WithVerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				response.RespondUnauthorized(w, ulanderrors.ErrTokenNotProvided)
				return
			}
			response.RespondBadRequest(w, ulanderrors.ErrTokenNotFound)
			return
		}

		tokenString := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTKey()), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response.RespondUnauthorized(w, ulanderrors.ErrTokenInvalidSignature)
				return
			}
			response.RespondBadRequest(w, ulanderrors.ErrTokenInvalidContent)
			return
		}

		if !token.Valid {
			response.RespondUnauthorized(w, ulanderrors.ErrTokenExpired)
			return
		}

		userRepo := GetUserRepository()
		user, err := userRepo.getUserById(claims.UserId)
		if err != nil {
			response.RespondBadRequest(w, ulanderrors.ErrTokenUserIdDoesNotExist)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
