package services

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const issuer = "OSTO CLI"

type TOTPService struct{}

func NewTOTPService() *TOTPService {
	return &TOTPService{}
}

func (s *TOTPService) Generate(username string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: username,
	})
}

func (s *TOTPService) Verify(secret, code string) bool {
	return totp.Validate(code, secret)
}

func (s *TOTPService) Secret(key *otp.Key) string {
	return key.Secret()
}

func (s *TOTPService) URL(key *otp.Key) string {
	return key.URL()
}