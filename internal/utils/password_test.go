package utils

import "testing"

// TestHashPassword verifies that a plain-text password
// is successfully hashed using bcrypt.
func TestHashPassword(t *testing.T) {

	password := "Aditya@123"

	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hash == "" {
		t.Fatal("expected hashed password, got empty string")
	}

	if hash == password {
		t.Fatal("hashed password should not match plain password")
	}
}

// TestCheckPasswordSuccess verifies that password comparison
// succeeds for the correct password.
func TestCheckPasswordSuccess(t *testing.T) {

	password := "Aditya@123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	err = CheckPassword(hash, password)

	if err != nil {
		t.Fatalf("expected password to match, got %v", err)
	}
}

// TestCheckPasswordFailure verifies that password comparison
// fails for an incorrect password.
func TestCheckPasswordFailure(t *testing.T) {

	password := "Aditya@123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	err = CheckPassword(hash, "WrongPassword@123")

	if err == nil {
		t.Fatal("expected password mismatch error")
	}
}