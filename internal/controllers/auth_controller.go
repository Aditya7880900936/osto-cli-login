package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Aditya7880900936/osto-cli-login/internal/services"
)

type AuthController struct {
	authService *services.AuthService
	reader      *bufio.Reader
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (c *AuthController) Register() {

	fmt.Print("Username: ")
	username, _ := c.reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := c.reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err := c.authService.Register(username, password)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Println("✅ User registered successfully.")
}