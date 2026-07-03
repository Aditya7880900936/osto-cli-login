package session

import (
	"time"

	"github.com/Aditya7880900936/osto-cli-login/internal/config"
	"github.com/Aditya7880900936/osto-cli-login/internal/models"
)

type SessionManager struct {
	currentUser *models.User
	expiresAt   time.Time
}

func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

func (s *SessionManager) Create(user *models.User) {
	s.currentUser = user
	s.expiresAt = time.Now().Add(config.SessionTimeout)
}

func (s *SessionManager) Destroy() {
	s.currentUser = nil
	s.expiresAt = time.Time{}
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

func (s *SessionManager) CurrentUser() *models.User {

	if !s.IsAuthenticated() {
		return nil
	}

	return s.currentUser
}

func (s *SessionManager) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s *SessionManager) Refresh() {
	if s.currentUser != nil {
		s.expiresAt = time.Now().Add(config.SessionTimeout)
	}
}