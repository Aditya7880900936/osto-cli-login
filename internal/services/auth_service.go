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

// AuthService contains the core business logic
// for authentication, registration, and MFA.
type AuthService struct {
	userRepo    *repository.UserRepository
	totpService *TOTPService
}

// NewAuthService creates a new AuthService instance.
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		totpService: NewTOTPService(),
	}
}

// Register validates user credentials,
// hashes the password, and creates a new account.
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

// Login authenticates a user,
// applies account lockout rules,
// and determines whether MFA verification is required.
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

// VerifyOTP validates the TOTP code
// for users with multi-factor authentication enabled.
func (s *AuthService) VerifyOTP(user *models.User, code string) error {

	if !user.MFAEnabled {
		return nil
	}

	if !s.totpService.Verify(user.TOTPSecret, code) {
		return errors.New("invalid otp")
	}

	return nil
}

// Generate2FASetup generates a TOTP secret
// and provisioning URL for Google Authenticator.
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

// Confirm2FA verifies the provided OTP
// and enables two-factor authentication.
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

// Disable2FA disables multi-factor authentication
// for the specified user.
func (s *AuthService) Disable2FA(username string) error {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	user.MFAEnabled = false
	user.TOTPSecret = ""

	return s.userRepo.Update(user)
}
