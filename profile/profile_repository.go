package profile

import (
	"userland/appcontext"
	"userland/auth"

	"github.com/jmoiron/sqlx"
)

const (
	UPDATE_PROFILE_BY_ID_QUERY = "UPDATE \"user\" SET fullname=$1, location=$2, bio=$3, web=$4 WHERE id=$5"
)

type profileRepository struct {
	db *sqlx.DB
}

func getProfileRepository() *profileRepository {
	repo := profileRepository{appcontext.GetDB()}
	return &repo
}

func (repo *profileRepository) updateUserProfile(user *auth.User, newUserProfile UserProfile) error {
	_, err := repo.db.Queryx(UPDATE_PROFILE_BY_ID_QUERY, newUserProfile.Fullname, newUserProfile.Location, newUserProfile.Bio, newUserProfile.Web, user.Id)
	return err
}
