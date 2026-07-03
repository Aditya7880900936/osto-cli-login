package session

import (
	"time"

	"github.com/Aditya7880900936/osto-cli-login/internal/models"
)

type SessionManager struct {
	currentUser *models.User
	expiresAt   time.Time
	timeout     time.Duration
}

func NewSessionManager(timeout time.Duration) *SessionManager {
	return &SessionManager{
		timeout: timeout,
	}
}

func (s *SessionManager) Create(user *models.User) {
	s.currentUser = user
	s.expiresAt = time.Now().Add(s.timeout)
}

func (s *SessionManager) Destroy() {
	s.currentUser = nil
	s.expiresAt = time.Time{}
}

func (s *SessionManager) CurrentUser() *models.User {
	if !s.IsAuthenticated() {
		return nil
	}
	return s.currentUser
}

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

func (s *SessionManager) ExpiresAt() time.Time {
	return s.expiresAt
}