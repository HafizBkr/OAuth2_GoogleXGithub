package user

import (
    "errors"
    "anonymdevs/models"
)

type UserService interface {
    RegisterUser(user *models.User) error
    GetUserByEmail(email string) (*models.User, error)
}

type DefaultUserService struct {
    Repo UserRepository
}

func (s *DefaultUserService) RegisterUser(user *models.User) error {
    if user == nil {
        return errors.New("user cannot be nil")
    }
    // Apply business logic before saving
    return s.Repo.SaveOrUpdateUser(user)
}

func (s *DefaultUserService) GetUserByEmail(email string) (*models.User, error) {
    return s.Repo.GetUserByEmail(email)
}
