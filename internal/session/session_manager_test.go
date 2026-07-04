package session

import (
	"testing"
	"time"

	"github.com/Aditya7880900936/osto-cli-login/internal/models"
)

// TestNewSessionManager verifies that a new session manager
// is initialized correctly.
func TestNewSessionManager(t *testing.T) {

	session := NewSessionManager()

	if session == nil {
		t.Fatal("expected session manager instance")
	}
}

// TestCreateSession verifies that a session
// is successfully created for a user.
func TestCreateSession(t *testing.T) {

	session := NewSessionManager()

	user := &models.User{
		Username: "aditya",
	}

	session.Create(user)

	if !session.IsAuthenticated() {
		t.Fatal("expected authenticated session")
	}

	if session.CurrentUser() == nil {
		t.Fatal("expected current user")
	}

	if session.CurrentUser().Username != "aditya" {
		t.Fatal("unexpected username")
	}
}

// TestDestroySession verifies that the current
// session is properly destroyed.
func TestDestroySession(t *testing.T) {

	session := NewSessionManager()

	user := &models.User{
		Username: "aditya",
	}

	session.Create(user)
	session.Destroy()

	if session.IsAuthenticated() {
		t.Fatal("expected session to be destroyed")
	}

	if session.CurrentUser() != nil {
		t.Fatal("expected nil user")
	}
}

// TestRefreshSession verifies that refreshing
// a session extends its expiration time.
func TestRefreshSession(t *testing.T) {

	session := NewSessionManager()

	user := &models.User{
		Username: "aditya",
	}

	session.Create(user)

	oldExpiry := session.ExpiresAt()

	time.Sleep(2 * time.Second)

	session.Refresh()

	if !session.ExpiresAt().After(oldExpiry) {
		t.Fatal("expected session expiry to be refreshed")
	}
}

// TestCurrentUser verifies that the current
// authenticated user is returned correctly.
func TestCurrentUser(t *testing.T) {

	session := NewSessionManager()

	user := &models.User{
		Username: "aditya",
	}

	session.Create(user)

	current := session.CurrentUser()

	if current == nil {
		t.Fatal("expected current user")
	}

	if current.Username != user.Username {
		t.Fatal("unexpected user returned")
	}
}