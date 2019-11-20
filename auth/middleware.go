package auth

import (
	"context"
	"net/http"
	"userland/config"
	ulanderrors "userland/errors"
	"userland/response"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	UserRepo userRepositoryInterface
}

func (middleware AuthMiddleware) WithVerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			log.Info(err)
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
			log.Info(err)
			if err == jwt.ErrSignatureInvalid {
				response.RespondUnauthorized(w, ulanderrors.ErrTokenInvalidSignature)
				return
			}
			valErr, ok := err.(*jwt.ValidationError)
			if ok {
				if (valErr.Errors & jwt.ValidationErrorSignatureInvalid) != 0 {
					response.RespondUnauthorized(w, ulanderrors.ErrTokenInvalidSignature)
					return
				}
				if (valErr.Errors & jwt.ValidationErrorExpired) != 0 {
					response.RespondUnauthorized(w, ulanderrors.ErrTokenExpired)
					return
				}
			}
			response.RespondBadRequest(w, ulanderrors.ErrTokenInvalidContent)
			return
		}

		if !token.Valid {
			log.Info("Token expired")
			response.RespondUnauthorized(w, ulanderrors.ErrTokenExpired)
			return
		}

		user, err := middleware.UserRepo.getUserById(claims.UserId)
		if err != nil {
			log.Warn(err)
			response.RespondBadRequest(w, ulanderrors.ErrTokenUserIdDoesNotExist)
			return
		}

		log.Info("Authentication successful")
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
