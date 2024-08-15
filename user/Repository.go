package user

import (
    "github.com/jmoiron/sqlx"
    "github.com/google/uuid"
    "anonymdevs/models"
)

type UserRepository interface {
    SaveOrUpdateUser(user *models.User) error
    GetUserByEmail(email string) (*models.User, error)
    GetRoleByLabel(label string) (*models.Role, error)
    AssignRoleToUser(userID uuid.UUID, roleID uuid.UUID) error
    GetUserByID(id string) (*models.User, error)
}

type UserRepo struct {
    db *sqlx.DB
}
func NewUserRepo(db *sqlx.DB) *UserRepo {
    return &UserRepo{
        db: db,
    }
}
func (r *UserRepo) SaveOrUpdateUser(user *models.User) error {
    if user.ID == "" {
        user.ID = uuid.New().String()
        query := `INSERT INTO users (id, username, email, auth_provider, profile_image, created_at, updated_at, is_validated)
                  VALUES (:id, :username, :email, :auth_provider, :profile_image, :created_at, :updated_at, :is_validated)`
        _, err := r.db.NamedExec(query, user)
        return err
    } else {
        query := `UPDATE users SET username=:username, auth_provider=:auth_provider, profile_image=:profile_image, updated_at=:updated_at, is_validated=:is_validated WHERE id=:id`
        _, err := r.db.NamedExec(query, user)
        return err
    }
}

func (r *UserRepo ) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    query := `SELECT id, username, email, auth_provider, profile_image, created_at, updated_at, is_validated FROM users WHERE email=$1`
    err := r.db.Get(&user, query, email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepo) GetRoleByLabel(label string) (*models.Role, error) {
    var role models.Role
    query := `SELECT id, label FROM roles WHERE label=$1`
    err := r.db.Get(&role, query, label)
    if err != nil {
        return nil, err
    }
    return &role, nil
}

func (r *UserRepo) AssignRoleToUser(userID uuid.UUID, roleID uuid.UUID) error {
    query := `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)`
    _, err := r.db.Exec(query, userID, roleID)
    return err
}

func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
    var user models.User
    query := `SELECT id, username, email, auth_provider, profile_image, created_at, updated_at, is_validated FROM users WHERE id=$1`
    err := r.db.Get(&user, query, id)
    if err != nil {
        return nil, err
    }
    return &user, nil
}