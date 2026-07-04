package services

import (
	"testing"
    "time"
	
	"github.com/pquerna/otp/totp"
)

// TestNewTOTPService verifies that a new
// TOTP service instance is created successfully.
func TestNewTOTPService(t *testing.T) {

	service := NewTOTPService()

	if service == nil {
		t.Fatal("expected TOTP service instance")
	}
}

// TestGenerate verifies that a TOTP secret
// and provisioning key are generated successfully.
func TestGenerate(t *testing.T) {

	service := NewTOTPService()

	key, err := service.Generate("aditya")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key == nil {
		t.Fatal("expected generated key")
	}

	if key.Secret() == "" {
		t.Fatal("expected non-empty secret")
	}

	if key.URL() == "" {
		t.Fatal("expected provisioning URL")
	}
}

// TestVerifySuccess verifies that a valid
// TOTP code is accepted.
func TestVerifySuccess(t *testing.T) {

	service := NewTOTPService()

	key, err := service.Generate("aditya")
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	code, err := totp.GenerateCode(key.Secret(), Now())
	if err != nil {
		t.Fatalf("failed to generate otp: %v", err)
	}

	if !service.Verify(key.Secret(), code) {
		t.Fatal("expected OTP verification to succeed")
	}
}

// TestVerifyFailure verifies that an
// invalid TOTP code is rejected.
func TestVerifyFailure(t *testing.T) {

	service := NewTOTPService()

	key, err := service.Generate("aditya")
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	if service.Verify(key.Secret(), "123456") {
		t.Fatal("expected OTP verification to fail")
	}
}

// Now returns the current time.
// Defined separately to simplify testing.
func Now() time.Time {
	return time.Now()
}