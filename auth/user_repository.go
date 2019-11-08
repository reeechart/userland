package auth

import (
	"userland/appcontext"

	"github.com/jmoiron/sqlx"
)

const (
	CREATE_USER_QUERY = "INSERT INTO \"user\"(fullname, email, password) VALUES ($1, $2, $3)"
)

type userRepositoryInterface interface {
	createNewUser(u User)
}

type userRepository struct {
	db *sqlx.DB
}

func getUserRepository() *userRepository {
	repo := userRepository{appcontext.GetDB()}
	return &repo
}

func (repo *userRepository) createNewUser(u userRegistration) error {
	hashedPassword := u.hashedPassword()
	_, err := repo.db.Queryx(CREATE_USER_QUERY, u.Fullname, u.Email, hashedPassword)
	return err
}
