package auth

import (
	"context"
	"net/http"
	"userland/config"
	"userland/response"

	"github.com/dgrijalva/jwt-go"
)

func WithVerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				response.RespondUnauthorized(w, TOKEN_NOT_PROVIDED, err)
				return
			}
			response.RespondBadRequest(w, TOKEN_CANNOT_BE_FOUND, err)
			return
		}

		tokenString := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTKey()), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response.RespondUnauthorized(w, TOKEN_INVALID_SIGNATURE, err)
				return
			}
			response.RespondBadRequest(w, TOKEN_INVALID_CONTENT, err)
			return
		}

		if !token.Valid {
			response.RespondUnauthorized(w, TOKEN_EXPIRED, err)
			return
		}

		userRepo := GetUserRepository()
		user, err := userRepo.getUserById(claims.UserId)
		if err != nil {
			response.RespondBadRequest(w, USER_ID_DOES_NOT_EXIST, err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
