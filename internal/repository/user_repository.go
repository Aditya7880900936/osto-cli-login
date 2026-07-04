package repository

import (
	"github.com/Aditya7880900936/osto-cli-login/internal/database"
	"github.com/Aditya7880900936/osto-cli-login/internal/models"
)

// UserRepository provides database operations
// related to user management.
type UserRepository struct{}

// NewUserRepository creates and returns
// a new UserRepository instance.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create inserts a new user into the database.
func (r *UserRepository) Create(user *models.User) error {
	return database.GetDB().Create(user).Error
}

// FindByUsername retrieves a user by username.
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

// Update persists changes made to an existing user.
func (r *UserRepository) Update(user *models.User) error {
	return database.GetDB().Save(user).Error
}

// IncrementFailedAttempts increases the failed login counter.
func (r *UserRepository) IncrementFailedAttempts(user *models.User) error {
	user.FailedAttempts++
	return database.GetDB().Save(user).Error
}

// ResetFailedAttempts clears failed login attempts
// and removes any active account lock.
func (r *UserRepository) ResetFailedAttempts(user *models.User) error {
	user.FailedAttempts = 0
	user.LockedUntil = nil
	return database.GetDB().Save(user).Error
}
