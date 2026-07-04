package utils

import (
	"errors"
	"regexp"
	"strings"
)

// Regular expressions used for username
// and password validation.
var (
	usernameRegex  = regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`)
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex = regexp.MustCompile(`[a-z]`)
	numberRegex    = regexp.MustCompile(`[0-9]`)
	specialRegex   = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
)

// ValidateUsername validates the username
// according to the application's rules.
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)

	if username == "" {
		return errors.New("username is required")
	}

	if !usernameRegex.MatchString(username) {
		return errors.New("username must be 3-30 characters and contain only letters, numbers and underscore")
	}

	return nil
}

// ValidatePassword validates the password
// against the defined security requirements.
func ValidatePassword(password string) error {

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if !uppercaseRegex.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !lowercaseRegex.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !numberRegex.MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	if !specialRegex.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
