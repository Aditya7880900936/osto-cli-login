package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Aditya7880900936/osto-cli-login/internal/controllers"
)

type CLI struct {
	authController *controllers.AuthController
	reader         *bufio.Reader
}

func NewCLI(authController *controllers.AuthController) *CLI {
	return &CLI{
		authController: authController,
		reader:         bufio.NewReader(os.Stdin),
	}
}

func (c *CLI) Start() {

	fmt.Println("===================================")
	fmt.Println("     OSTO CLI Authentication")
	fmt.Println("===================================")
	fmt.Println("Type 'help' to see available commands.")

	for {
		fmt.Print("\n> ")

		command, _ := c.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {

		case "register":
			c.authController.Register()

		case "login":
            c.authController.Login()

		case "whoami":
	        c.authController.WhoAmI()
		
		case "logout":
	        c.authController.Logout()

		case "help":
			c.printHelp()

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Type 'help'.")
		}
	}
}

func (c *CLI) printHelp() {
	fmt.Println("\nAvailable Commands")
	fmt.Println("------------------")
	fmt.Println("register")
	fmt.Println("login")
	fmt.Println("whoami")
	fmt.Println("logout")
	fmt.Println("help")
	fmt.Println("exit")
}