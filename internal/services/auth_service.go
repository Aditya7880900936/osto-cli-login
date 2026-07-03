package services

import (
	"errors"
	"time"

	"github.com/Aditya7880900936/osto-cli-login/internal/config"
	"github.com/Aditya7880900936/osto-cli-login/internal/models"
	"github.com/Aditya7880900936/osto-cli-login/internal/repository"
	"github.com/Aditya7880900936/osto-cli-login/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	totpService *TOTPService
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		totpService: NewTOTPService(),
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

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: hash,
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(username, password string) (*models.User, bool, error) {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, errors.New("invalid username or password")
		}
		return nil, false, err
	}

	if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
		return nil, false, errors.New("account is temporarily locked")
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {

		user.FailedAttempts++

		if user.FailedAttempts >= config.MaxFailedAttempts {

			lockUntil := time.Now().Add(config.LockDuration)

			user.LockedUntil = &lockUntil
		}

		if err := s.userRepo.Update(user); err != nil {
			return nil, false, err
		}

		if user.FailedAttempts >= config.MaxFailedAttempts {
			return nil, false, errors.New("account locked for 5 minutes")
		}

		return nil, false, errors.New("invalid username or password")
	}

	user.FailedAttempts = 0
	user.LockedUntil = nil

	now := time.Now()
	user.LastLogin = &now

	if err := s.userRepo.Update(user); err != nil {
		return nil, false, err
	}

	if user.MFAEnabled {
		return user, true, nil
	}

	return user, false, nil
}

func (s *AuthService) VerifyOTP(user *models.User, code string) error {

	if !user.MFAEnabled {
		return nil
	}

	if !s.totpService.Verify(user.TOTPSecret, code) {
		return errors.New("invalid otp")
	}

	return nil
}

func (s *AuthService) Generate2FASetup(username string) (string, string, error) {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", "", err
	}

	if user.MFAEnabled {
		return "", "", errors.New("2FA is already enabled")
	}

	key, err := s.totpService.Generate(user.Username)
	if err != nil {
		return "", "", err
	}

	return key.Secret(), key.URL(), nil
}

func (s *AuthService) Confirm2FA(username, secret, code string) error {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	if !s.totpService.Verify(secret, code) {
		return errors.New("invalid otp")
	}

	user.TOTPSecret = secret
	user.MFAEnabled = true

	return s.userRepo.Update(user)
}

func (s *AuthService) Disable2FA(username string) error {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	user.MFAEnabled = false
	user.TOTPSecret = ""

	return s.userRepo.Update(user)
}