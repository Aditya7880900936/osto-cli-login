package utils

import "testing"

func TestValidateUsername(t *testing.T) {

	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{
			name:     "valid username",
			username: "aditya123",
			wantErr:  false,
		},
		{
			name:     "empty username",
			username: "",
			wantErr:  true,
		},
		{
			name:     "username too short",
			username: "ab",
			wantErr:  true,
		},
		{
			name:     "invalid characters",
			username: "adi@123",
			wantErr:  true,
		},
		{
			name:     "underscore allowed",
			username: "adi_test",
			wantErr:  false,
		},
	}

	for _, tt := range tests {

		err := ValidateUsername(tt.username)

		if (err != nil) != tt.wantErr {
			t.Errorf("%s: expected error=%v, got=%v", tt.name, tt.wantErr, err)
		}
	}
}

func TestValidatePassword(t *testing.T) {

	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "Aditya@123",
			wantErr:  false,
		},
		{
			name:     "too short",
			password: "Ad@12",
			wantErr:  true,
		},
		{
			name:     "missing uppercase",
			password: "aditya@123",
			wantErr:  true,
		},
		{
			name:     "missing lowercase",
			password: "ADITYA@123",
			wantErr:  true,
		},
		{
			name:     "missing number",
			password: "Aditya@Test",
			wantErr:  true,
		},
		{
			name:     "missing special character",
			password: "Aditya123",
			wantErr:  true,
		},
	}

	for _, tt := range tests {

		err := ValidatePassword(tt.password)

		if (err != nil) != tt.wantErr {
			t.Errorf("%s: expected error=%v, got=%v", tt.name, tt.wantErr, err)
		}
	}
}