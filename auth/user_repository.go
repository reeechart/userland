package auth

import (
	"errors"
	"userland/appcontext"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	CREATE_USER_QUERY                     = "INSERT INTO \"user\" (fullname, email, password, verification_token) VALUES ($1, $2, $3, $4)"
	SELECT_USER_BY_EMAIL_QUERY            = "SELECT * FROM \"user\" WHERE email=$1"
	UPDATE_VERIF_TOKEN_QUERY              = "UPDATE \"user\" SET verification_token=$1 WHERE id=$2"
	UPDATE_RESET_PASS_TOKEN_QUERY         = "UPDATE \"user\" SET reset_password_token=$1 WHERE id=$2"
	SELECT_USER_BY_RESET_PASS_TOKEN_QUERY = "SELECT * FROM \"user\" WHERE reset_password_token=$1"
	RESET_PASSWORD_QUERY                  = "UPDATE \"user\" SET password=$1, reset_password_token=NULL WHERE id=$2"
	SELECT_USER_BY_ID_QUERY               = "SELECT * FROM \"user\" WHERE id=$1"
	UPDATE_VERIFIED_QUERY                 = "UPDATE \"user\" SET verification_token=NULL, verified=true WHERE id=$1"
)

type userRepositoryInterface interface {
	createNewUser(user userRegistration) error
	verifyUser(recipient string, token string) error
	loginUser(email string, password string) error
	forgetPassword(email string) error
	getUserByEmail(email string) (*User, error)
	resetPassword(token string, password string) error
	getUserByResetPasswordToken(token string) (*User, error)
	getUserById(id int) (*User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func GetUserRepository() *userRepository {
	repo := userRepository{appcontext.GetDB()}
	return &repo
}

func (repo *userRepository) createNewUser(user userRegistration) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	stmt, err := repo.db.Preparex(CREATE_USER_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(user.Fullname, user.Email, string(passwordHash), generateToken())
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

	stmt, err := repo.db.Preparex(UPDATE_VERIFIED_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(user.Id)
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
	stmt, err := repo.db.Preparex(UPDATE_RESET_PASS_TOKEN_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(generateToken(), user.Id)
	return err
}

func (repo *userRepository) getUserByEmail(email string) (*User, error) {
	stmt, err := repo.db.Preparex(SELECT_USER_BY_EMAIL_QUERY)
	if err != nil {
		return nil, err
	}
	row, err := stmt.Queryx(email)
	if err != nil {
		return nil, err
	}
	var user User
	row.Next()
	defer row.Close()
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

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	stmt, err := repo.db.Preparex(RESET_PASSWORD_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(passwordHash, user.Id)
	return err
}

func (repo *userRepository) getUserByResetPasswordToken(token string) (*User, error) {
	stmt, err := repo.db.Preparex(SELECT_USER_BY_RESET_PASS_TOKEN_QUERY)
	if err != nil {
		return nil, err
	}
	row, err := stmt.Queryx(token)
	if err != nil {
		return nil, err
	}
	var user User
	row.Next()
	defer row.Close()
	err = row.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) getUserById(id int) (*User, error) {
	stmt, err := repo.db.Preparex(SELECT_USER_BY_ID_QUERY)
	if err != nil {
		return nil, err
	}
	row, err := stmt.Queryx(SELECT_USER_BY_ID_QUERY, id)
	if err != nil {
		return nil, err
	}
	var user User
	row.Next()
	defer row.Close()
	err = row.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
