package auth

import (
	"errors"
	"userland/appcontext"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	CREATE_USER_QUERY                     = "INSERT INTO \"user\" (fullname, email, password) VALUES ($1, $2, $3) RETURNING id"
	SELECT_USER_BY_EMAIL_QUERY            = "SELECT * FROM \"user\" WHERE email=$1"
	UPDATE_VERIF_TOKEN_QUERY              = "UPDATE \"user\" SET verification_token=$1 WHERE id=$2"
	UPDATE_RESET_PASS_TOKEN_QUERY         = "UPDATE \"user\" SET reset_password_token=$1 WHERE id=$2"
	SELECT_USER_BY_RESET_PASS_TOKEN_QUERY = "SELECT * FROM \"user\" WHERE reset_password_token=$1"
	UPDATE_PASSWORD_QUERY                 = "UPDATE \"user\" SET password=$1 WHERE id=$2"
	SELECT_USER_BY_ID_QUERY               = "SELECT * FROM \"user\" WHERE id=$1"
	DELETE_RESET_PASS_TOKEN_QUERY         = "UPDATE \"user\" SET reset_password_token=NULL WHERE id=$1"
	UPDATE_VERIFIED_QUERY                 = "UPDATE \"user\" SET verification_token=NULL, verified=true WHERE id=$1"
)

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

	var (
		tx        *sqlx.Tx
		newUserId int
	)

	tx, err = repo.db.Beginx()
	err = tx.Get(&newUserId, CREATE_USER_QUERY, user.Fullname, user.Email, string(passwordHash))
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(UPDATE_VERIF_TOKEN_QUERY, generateToken(), newUserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (repo *userRepository) verifyUser(recipient string, token string) error {
	user, err := repo.getUserByEmail(recipient)
	if err != nil {
		return err
	}

	if user.VerificationToken.String != token {
		return errors.New("Tokens don't match")
	}

	_, err = repo.db.Queryx(UPDATE_VERIFIED_QUERY, user.Id)
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

func (repo *userRepository) forgetPassword(email string) error {
	user, err := repo.getUserByEmail(email)
	if err != nil {
		return err
	}
	_, err = repo.db.Queryx(UPDATE_RESET_PASS_TOKEN_QUERY, generateToken(), user.Id)
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

func (repo *userRepository) resetPassword(token string, password string) error {
	user, err := repo.getUserByResetPasswordToken(token)
	if err != nil {
		return err
	}
	var tx *sqlx.Tx
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	tx, err = repo.db.Beginx()
	_, err = tx.Exec(UPDATE_PASSWORD_QUERY, passwordHash, user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(DELETE_RESET_PASS_TOKEN_QUERY, user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (repo *userRepository) getUserByResetPasswordToken(token string) (*User, error) {
	row, err := repo.db.Queryx(SELECT_USER_BY_RESET_PASS_TOKEN_QUERY, token)
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

func (repo *userRepository) getUserById(id int) (*User, error) {
	row, err := repo.db.Queryx(SELECT_USER_BY_ID_QUERY, id)
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
