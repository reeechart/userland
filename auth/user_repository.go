package auth

import (
	"userland/appcontext"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	CREATE_USER_QUERY          = "INSERT INTO \"user\" (fullname, email, password) VALUES ($1, $2, $3)"
	SELECT_USER_BY_EMAIL_QUERY = "SELECT * FROM \"user\" WHERE email=$1"
	UPDATE_VERIF_TOKEN_QUERY   = "UPDATE \"user\" SET verification_token=$1 WHERE id=$2"
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

func (repo *userRepository) verifyUser(recipient string) error {
	user, err := repo.getUserByEmail(recipient)
	if err != nil {
		return err
	}
	_, err = repo.db.Queryx(UPDATE_VERIF_TOKEN_QUERY, generateToken(), user.Id)
	return err
}

func (repo *userRepository) loginUser(email string, password string) error {
	user, err := repo.getUserByEmail(email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}

func (repo *userRepository) getUserByEmail(email string) (*User, error) {
	row, err := repo.db.Queryx(SELECT_USER_BY_EMAIL_QUERY, email)
	if err != nil {
		return nil, err
	}
	var user User
	row.Next()
	err = row.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
