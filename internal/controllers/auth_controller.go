package controllers

import (
	"fmt"

	"github.com/Aditya7880900936/osto-cli-login/internal/services"
	"github.com/Aditya7880900936/osto-cli-login/internal/session"
	"github.com/Aditya7880900936/osto-cli-login/internal/utils"
)

// AuthController handles all user interactions
// related to authentication and session management.
type AuthController struct {
	authService    *services.AuthService
	sessionManager *session.SessionManager
}

// NewAuthController creates and returns
// a new authentication controller.
func NewAuthController(
	authService *services.AuthService,
	sessionManager *session.SessionManager,
) *AuthController {
	return &AuthController{
		authService:    authService,
		sessionManager: sessionManager,
	}
}

// Register collects user credentials
// and creates a new user account.
func (c *AuthController) Register() {

	username := utils.ReadLine("Username: ")
	password := utils.ReadPassword("Password: ")

	if err := c.authService.Register(username, password); err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Println("✅ User registered successfully.")
}

// Login authenticates the user and
// performs TOTP verification if enabled.
func (c *AuthController) Login() {

	username := utils.ReadLine("Username: ")
	password := utils.ReadPassword("Password: ")

	user, requiresOTP, err := c.authService.Login(username, password)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	if requiresOTP {

		otp := utils.ReadLine("OTP: ")

		if err := c.authService.VerifyOTP(user, otp); err != nil {
			fmt.Println("❌", err)
			return
		}
	}

	c.sessionManager.Create(user)

	fmt.Println("✅ Login successful.")

	c.WhoAmI()
}

// Enable2FA generates a TOTP secret,
// verifies the OTP, and enables multi-factor authentication.
func (c *AuthController) Enable2FA() {

	if !c.sessionManager.IsAuthenticated() {
		fmt.Println("❌ Login required.")
		return
	}

	user := c.sessionManager.CurrentUser()

	secret, url, err := c.authService.Generate2FASetup(user.Username)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Println("\n====== Google Authenticator ======")
	fmt.Println("Secret :", secret)
	fmt.Println("OTPAuth URL:")
	fmt.Println(url)
	fmt.Println("==================================")

	code := utils.ReadLine("Enter OTP: ")

	if err := c.authService.Confirm2FA(user.Username, secret, code); err != nil {
		fmt.Println("❌", err)
		return
	}

	user.MFAEnabled = true
	user.TOTPSecret = secret

	fmt.Println("✅ 2FA Enabled Successfully.")
}

// Disable2FA disables two-factor authentication
// for the currently authenticated user.
func (c *AuthController) Disable2FA() {

	if !c.sessionManager.IsAuthenticated() {
		fmt.Println("❌ Login required.")
		return
	}

	user := c.sessionManager.CurrentUser()

	if err := c.authService.Disable2FA(user.Username); err != nil {
		fmt.Println("❌", err)
		return
	}

	user.MFAEnabled = false
	user.TOTPSecret = ""

	fmt.Println("✅ 2FA Disabled Successfully.")
}

// WhoAmI displays information about
// the currently authenticated user.
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

// Logout terminates the current user session.
func (c *AuthController) Logout() {

	if !c.sessionManager.IsAuthenticated() {
		fmt.Println("❌ You are not logged in.")
		return
	}

	c.sessionManager.Destroy()

	fmt.Println("✅ Logged out successfully.")
}
