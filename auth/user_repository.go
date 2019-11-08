package auth

import (
	"userland/appcontext"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	CREATE_USER_QUERY = "INSERT INTO \"user\"(fullname, email, password) VALUES ($1, $2, $3)"
)

type userRepositoryInterface interface {
	createNewUser(user userRegistration)
	loginUser(user User)
}

type userRepository struct {
	db *sqlx.DB
}

func getUserRepository() *userRepository {
	repo := userRepository{appcontext.GetDB()}
	return &repo
}

func (repo *userRepository) createNewUser(user userRegistration) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = repo.db.Queryx(CREATE_USER_QUERY, user.Fullname, user.Email, string(passwordHash))
	return err
}
