package user

import (
    "anonymdevs/models"
    "errors"
    "time"
    "github.com/google/uuid"
)

type UserService interface {
    RegisterUser(userPayload *UserPayload, authProvider string) (*models.User, error)
    GetUserByEmail(email string) (*models.User, error)
    UpdateUser(user *models.User) error
    GetUserById(id string) (*models.User, error)
}

type DefaultUserService struct {
    Repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
    return &DefaultUserService{
        Repo: repo,
    }
}

func (s *DefaultUserService) RegisterUser(userPayload *UserPayload, authProvider string) (*models.User, error) {
    var roleLabel string
    if authProvider == "google" {
        roleLabel = "lecteur"
    } else if authProvider == "github" {
        roleLabel = "redacteur"
    } else {
        return nil, errors.New("unsupported auth provider")
    }

    userID := uuid.New()

    user := &models.User{
        ID:           userID.String(),
        Username:     userPayload.Username,
        Email:        userPayload.Email,
        AuthProvider: authProvider,
        ProfileImage: userPayload.ProfileImage,
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
        IsValidated:  false,
    }

    err := s.Repo.SaveOrUpdateUser(user)
    if err != nil {
        return nil, err
    }

    role, err := s.Repo.GetRoleByLabel(roleLabel)
    if err != nil {
        return nil, err
    }

    err = s.Repo.AssignRoleToUser(userID, role.ID)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *DefaultUserService) GetUserByEmail(email string) (*models.User, error) {
    return s.Repo.GetUserByEmail(email)
}
func (s *DefaultUserService) GetUserById(id string) (*models.User, error) {
    return s.Repo.GetUserByID(id)
}

func (s *DefaultUserService) UpdateUser(user *models.User) error {
    return s.Repo.SaveOrUpdateUser(user)
}
