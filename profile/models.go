package profile

import (
	"time"
)

type UserProfile struct {
	Id             int       `json:"id"`
	Fullname       string    `json:"fullname"`
	Location       string    `json:"location"`
	Bio            string    `json:"bio"`
	Web            string    `json:"web"`
	ProfilePicture []byte    `json:"picture" db:"picture"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
