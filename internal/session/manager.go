package session

import (
	"time"

	"github.com/Aditya7880900936/osto-cli-login/internal/config"
	"github.com/Aditya7880900936/osto-cli-login/internal/models"
)

// SessionManager manages the authenticated user session
// and tracks its expiration time.
type SessionManager struct {
	currentUser *models.User
	expiresAt   time.Time
}

// NewSessionManager creates and returns
// a new session manager instance.
func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

// Create starts a new authenticated session
// for the specified user.
func (s *SessionManager) Create(user *models.User) {
	s.currentUser = user
	s.expiresAt = time.Now().Add(config.SessionTimeout)
}

// Destroy clears the current authenticated session.
func (s *SessionManager) Destroy() {
	s.currentUser = nil
	s.expiresAt = time.Time{}
}

// IsAuthenticated checks whether
// the current session is still valid.
func (s *SessionManager) IsAuthenticated() bool {

	if s.currentUser == nil {
		return false
	}

	if time.Now().After(s.expiresAt) {
		s.Destroy()
		return false
	}

	return true
}

// CurrentUser returns the authenticated user
// if the current session is valid.
func (s *SessionManager) CurrentUser() *models.User {

	if !s.IsAuthenticated() {
		return nil
	}

	return s.currentUser
}

// ExpiresAt returns the session expiration time.
func (s *SessionManager) ExpiresAt() time.Time {
	return s.expiresAt
}

// Refresh extends the session expiration time
// for the currently authenticated user.
func (s *SessionManager) Refresh() {
	if s.currentUser != nil {
		s.expiresAt = time.Now().Add(config.SessionTimeout)
	}
}
