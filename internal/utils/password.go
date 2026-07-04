package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain-text password
// using the bcrypt algorithm.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CheckPassword compares a plain-text password
// against its bcrypt hash.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
