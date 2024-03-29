package profile

import (
	"userland/appcontext"
	"userland/auth"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	UPDATE_PROFILE_BY_ID_QUERY         = "UPDATE \"user\" SET fullname=$1, location=$2, bio=$3, web=$4 WHERE id=$5"
	CHANGE_EMAIL_BY_ID_QUERY           = "UPDATE \"user\" SET email=$1 WHERE id=$2"
	CHANGE_PASSWORD_BY_ID_QUERY        = "UPDATE \"user\" SET password=$1 WHERE id=$2"
	DELETE_USER_BY_ID_QUERY            = "DELETE FROM \"user\" WHERE id=$1"
	UPDATE_PROFILE_PICTURE_BY_ID_QUERY = "UPDATE \"user\" SET picture=$1 WHERE id=$2"
	DELETE_PROFILE_PICTURE_BY_ID_QUERY = "UPDATE \"user\" SET picture=NULL WHERE id=$1"
)

type profileRepositoryInterface interface {
	updateUserProfile(user *auth.User, newUserProfile UserProfile) error
	changeUserEmail(user *auth.User, newEmail string) error
	changeUserPassword(user *auth.User, oldPassword string, newPassword string) error
	deleteUser(user *auth.User, password string) error
	updateUserPicture(user *auth.User, picture []byte) error
	deleteUserPicture(user *auth.User) error
}

type profileRepository struct {
	db *sqlx.DB
}

func GetProfileRepository() *profileRepository {
	repo := profileRepository{appcontext.GetDB()}
	return &repo
}

func (repo *profileRepository) updateUserProfile(user *auth.User, newUserProfile UserProfile) error {
	stmt, err := repo.db.Preparex(UPDATE_PROFILE_BY_ID_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(newUserProfile.Fullname, newUserProfile.Location, newUserProfile.Bio, newUserProfile.Web, user.Id)
	return err
}

func (repo *profileRepository) changeUserEmail(user *auth.User, newEmail string) error {
	stmt, err := repo.db.Preparex(CHANGE_EMAIL_BY_ID_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(newEmail, user.Id)
	return err
}

func (repo *profileRepository) changeUserPassword(user *auth.User, oldPassword string, newPassword string) error {
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))

	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}

	stmt, err := repo.db.Preparex(CHANGE_PASSWORD_BY_ID_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(passwordHash, user.Id)
	return err
}

func (repo *profileRepository) deleteUser(user *auth.User, password string) error {
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	stmt, err := repo.db.Preparex(DELETE_USER_BY_ID_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(user.Id)
	return err
}

func (repo *profileRepository) updateUserPicture(user *auth.User, picture []byte) error {
	stmt, err := repo.db.Preparex(UPDATE_PROFILE_PICTURE_BY_ID_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(picture, user.Id)
	return err
}

func (repo *profileRepository) deleteUserPicture(user *auth.User) error {
	stmt, err := repo.db.Preparex(DELETE_PROFILE_PICTURE_BY_ID_QUERY)
	if err != nil {
		return err
	}
	_, err = stmt.Queryx(user.Id)
	return err
}
