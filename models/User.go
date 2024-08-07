package models

import (
    "time"
)


type User struct {
    ID           string    `json:"id" db:"id"`
    Username     string    `json:"username" db:"username"`
    Email        string    `json:"email" db:"email"`
    AuthProvider string    `json:"auth_provider" db:"auth_provider"`
    ProfileImage string    `json:"profile_image" db:"profile_image"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
    IsValidated  bool      `json:"is_validated" db:"is_validated"`
}


type LoggedInUser struct {
	User
	Roles []Role `json:"roles"`
}

func (u *LoggedInUser) HasRole(role string) bool {
	ok := false
	for _, r := range u.Roles {
		if r.Label == role {
			ok = true
			break
		}
	}
	return ok
}

