package models

import(
	"time"
)

type User struct {
    ID            string    `json:"id"`
    Username      string    `json:"username"`
    ProfilePicture string    `json:"profile_picture"`
    Email         string    `json:"email"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
