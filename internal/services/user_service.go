package services

import (
	"chat_app/internal/models"
	"chat_app/internal/repositories"
	"errors"
)

var (
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
	ErrUserNotFound             = errors.New("user not found")
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (us *UserService) CreateUser(user *models.User) error {
	existingUser, err := us.UserRepo.GetUserByPhoneNumber(user.PhoneNumber)
	if err != nil && err != ErrUserNotFound {
		return err
	}
	if existingUser != nil {
		return ErrPhoneNumberAlreadyExists
	}

	err = us.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByID(userID int) (*models.User, error) {
	user, err := us.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
