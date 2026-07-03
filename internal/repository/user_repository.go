package repository

import (
	"github.com/Aditya7880900936/osto-cli-login/internal/database"
	"github.com/Aditya7880900936/osto-cli-login/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.User) error {
	return database.GetDB().Create(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	err := database.GetDB().
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return database.GetDB().Save(user).Error
}

func (r *UserRepository) IncrementFailedAttempts(user *models.User) error {
	user.FailedAttempts++
	return database.GetDB().Save(user).Error
}

func (r *UserRepository) ResetFailedAttempts(user *models.User) error {
	user.FailedAttempts = 0
	user.LockedUntil = nil
	return database.GetDB().Save(user).Error
}