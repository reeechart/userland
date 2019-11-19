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
			if err == http.ErrNoCookie {
				log.Info(err)
				response.RespondUnauthorized(w, ulanderrors.ErrTokenNotProvided)
				return
			}
			log.Info(err)
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
				log.Info(err)
				response.RespondUnauthorized(w, ulanderrors.ErrTokenInvalidSignature)
				return
			}
			log.Info(err)
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
