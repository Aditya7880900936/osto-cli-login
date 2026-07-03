package services

import (
	"errors"
	"time"

	"github.com/Aditya7880900936/osto-cli-login/internal/models"
	"github.com/Aditya7880900936/osto-cli-login/internal/repository"
	"github.com/Aditya7880900936/osto-cli-login/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(username, password string) error {

	if err := utils.ValidateUsername(username); err != nil {
		return err
	}

	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	_, err := s.userRepo.FindByUsername(username)

	if err == nil {
		return errors.New("username already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(username, password string) error {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("invalid username or password")
		}
		return err
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return errors.New("invalid username or password")
	}

	now := time.Now()
	user.LastLogin = &now

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	return nil
}