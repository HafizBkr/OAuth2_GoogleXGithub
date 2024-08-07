package models

type Role struct {
	ID    int    `json:"id" db:"id"`
	Label string `json:"label" db:"label"`
}
