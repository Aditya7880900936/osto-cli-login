package controllers

import (
	"fmt"

	"github.com/Aditya7880900936/osto-cli-login/internal/services"
	"github.com/Aditya7880900936/osto-cli-login/internal/session"
	"github.com/Aditya7880900936/osto-cli-login/internal/utils"
)

type AuthController struct {
	authService    *services.AuthService
	sessionManager *session.SessionManager
}

func NewAuthController(
	authService *services.AuthService,
	sessionManager *session.SessionManager,
) *AuthController {
	return &AuthController{
		authService:    authService,
		sessionManager: sessionManager,
	}
}

func (c *AuthController) Register() {

	username := utils.ReadLine("Username: ")
	password := utils.ReadPassword("Password: ")

	if err := c.authService.Register(username, password); err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Println("✅ User registered successfully.")
}

func (c *AuthController) Login() {

	username := utils.ReadLine("Username: ")
	password := utils.ReadPassword("Password: ")

	user, err := c.authService.Login(username, password)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	c.sessionManager.Create(user)

	fmt.Println("✅ Login successful.")

	c.WhoAmI()
}

func (c *AuthController) WhoAmI() {

	if !c.sessionManager.IsAuthenticated() {
		fmt.Println("❌ You are not logged in.")
		return
	}

	c.sessionManager.Refresh()

	user := c.sessionManager.CurrentUser()

	fmt.Println("\n========== USER ==========")
	fmt.Printf("Username          : %s\n", user.Username)
	fmt.Printf("Registered On     : %s\n", user.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("MFA Enabled       : %t\n", user.MFAEnabled)

	if user.LastLogin != nil {
		fmt.Printf("Last Login        : %s\n", user.LastLogin.Format("2006-01-02 15:04:05"))
	}

	fmt.Printf("Session Expires   : %s\n", c.sessionManager.ExpiresAt().Format("2006-01-02 15:04:05"))
	fmt.Println("==========================")
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