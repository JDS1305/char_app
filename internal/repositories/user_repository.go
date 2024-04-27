package repositories

import (
	"chat_app/internal/models"
	"database/sql"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	_, err := ur.DB.Exec("INSERT INTO users (name, phone_number) VALUES ($1, $2)", user.Name, user.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	row := ur.DB.QueryRow("SELECT id, name, phone_number FROM users WHERE id = $1", userID)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	row := ur.DB.QueryRow("SELECT id, name, phone_number FROM users WHERE phone_number = $1", phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
