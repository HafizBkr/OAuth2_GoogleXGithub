package user

import (
    "github.com/jmoiron/sqlx"
    "anonymdevs/models"
)

type UserRepository interface {
    SaveOrUpdateUser(user *models.User) error
    GetUserByEmail(email string) (*models.User, error)
}

type PostgresUserRepository struct {
    DB *sqlx.DB
}

func (r *PostgresUserRepository) SaveOrUpdateUser(user *models.User) error {
    var query string
    if user.ID == "" {
        query = `
            INSERT INTO users (username, email, auth_provider, profile_image, created_at, updated_at, is_validated)
            VALUES (:username, :email, :auth_provider, :profile_image, :created_at, :updated_at, :is_validated)
            RETURNING id
        `
        stmt, err := r.DB.PrepareNamed(query)
        if err != nil {
            return err
        }
        err = stmt.QueryRowx(user).Scan(&user.ID)
        if err != nil {
            return err
        }
    } else {
        query = `
            UPDATE users
            SET username = :username, auth_provider = :auth_provider, profile_image = :profile_image, updated_at = :updated_at, is_validated = :is_validated
            WHERE id = :id
        `
        _, err := r.DB.NamedExec(query, user)
        return err
    }
    return nil
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    query := `
        SELECT id, username, email, auth_provider, profile_image, created_at, updated_at, is_validated
        FROM users
        WHERE email = $1
    `
    err := r.DB.Get(&user, query, email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
