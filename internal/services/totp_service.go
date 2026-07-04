package services

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// issuer identifies the application in authenticator apps.
const issuer = "OSTO CLI"

// TOTPService provides utilities
// for generating and validating TOTP codes.
type TOTPService struct{}

// NewTOTPService creates a new TOTP service instance.
func NewTOTPService() *TOTPService {
	return &TOTPService{}
}

// Generate creates a new TOTP secret
// and provisioning key for the specified user.
func (s *TOTPService) Generate(username string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: username,
	})
}

// Verify validates a TOTP code
// using the provided secret.
func (s *TOTPService) Verify(secret, code string) bool {
	return totp.Validate(code, secret)
}

// Secret returns the shared secret
// from a generated TOTP key.
func (s *TOTPService) Secret(key *otp.Key) string {
	return key.Secret()
}

// URL returns the OTPAuth provisioning URL
// for configuring authenticator applications.
func (s *TOTPService) URL(key *otp.Key) string {
	return key.URL()
}
