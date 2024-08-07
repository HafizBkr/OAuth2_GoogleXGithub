package models

import (
    "time"
)

type SocialLink struct {
    ID        string `json:"id" db:"id"`
    UserID    string `json:"user_id" db:"user_id"`
    Platform  string    `json:"platform" db:"platform"`
    URL       string    `json:"url" db:"url"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
