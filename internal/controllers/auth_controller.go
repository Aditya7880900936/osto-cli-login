package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Aditya7880900936/osto-cli-login/internal/services"
	"github.com/Aditya7880900936/osto-cli-login/internal/session"
	"github.com/Aditya7880900936/osto-cli-login/internal/utils"
)

type AuthController struct {
	authService    *services.AuthService
	sessionManager *session.SessionManager
	reader         *bufio.Reader
}

func NewAuthController(
	authService *services.AuthService,
	sessionManager *session.SessionManager,
) *AuthController {

	return &AuthController{
		authService:    authService,
		sessionManager: sessionManager,
		reader:         bufio.NewReader(os.Stdin),
	}
}

func (c *AuthController) Register() {

	username := utils.ReadLine("Username: ")
	username = strings.TrimSpace(username)

	password := utils.ReadLine("Password: ")
	password = strings.TrimSpace(password)

	if err := c.authService.Register(username, password); err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Println("✅ User registered successfully.")
}

func (c *AuthController) Login() {

	username := utils.ReadLine("Username: ")
	username = strings.TrimSpace(username)

	password := utils.ReadPassword("Password: ")
	password = strings.TrimSpace(password)

	user, err := c.authService.Login(username, password)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	c.sessionManager.Create(user)

	fmt.Println("✅ Login successful.")
}

func (c *AuthController) WhoAmI() {

	if !c.sessionManager.IsAuthenticated() {
		fmt.Println("❌ You are not logged in.")
		return
	}

	user := c.sessionManager.CurrentUser()

	fmt.Println("\nCurrent User")
	fmt.Println("------------")
	fmt.Println("Username :", user.Username)
	fmt.Println("MFA      :", user.MFAEnabled)

	if user.LastLogin != nil {
		fmt.Println("Last Login:", user.LastLogin.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("Session Expires:", c.sessionManager.ExpiresAt().Format("2006-01-02 15:04:05"))
}

func (c *AuthController) Logout() {

	if !c.sessionManager.IsAuthenticated() {
		fmt.Println("❌ You are not logged in.")
		return
	}

	c.sessionManager.Destroy()

	fmt.Println("✅ Logged out successfully.")
}

func (c *AuthController) Enable2FA() {
	fmt.Println("Coming soon...")
}

func (c *AuthController) Disable2FA() {
	fmt.Println("Coming soon...")
}