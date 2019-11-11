package profile

import (
	"database/sql"
	"time"
)

type UserProfile struct {
	Id             int            `json:"id"`
	Fullname       string         `json:"fullname"`
	Location       sql.NullString `json:"location"`
	Bio            sql.NullString `json:"bio"`
	Web            sql.NullString `json:"web"`
	ProfilePicture []byte         `json:"picture" db:"picture"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
}
