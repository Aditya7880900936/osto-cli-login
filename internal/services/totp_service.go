package services

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TOTPService struct{}

func NewTOTPService() *TOTPService {
	return &TOTPService{}
}

func (s *TOTPService) Generate(username string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      "OSTO CLI",
		AccountName: username,
	})
}

func (s *TOTPService) Validate(secret, code string) bool {
	return totp.Validate(code, secret)
}