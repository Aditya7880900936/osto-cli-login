package cli

import (
	"fmt"

	"github.com/Aditya7880900936/osto-cli-login/internal/controllers"
	"github.com/chzyer/readline"
)

type CLI struct {
	authController *controllers.AuthController
	rl             *readline.Instance
}

func NewCLI(authController *controllers.AuthController) *CLI {

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "/tmp/osto_cli_history.tmp",
		AutoComplete:    Completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		panic(err)
	}

	return &CLI{
		authController: authController,
		rl:             rl,
	}
}

func (c *CLI) Start() {

	defer c.rl.Close()

	fmt.Println("===================================")
	fmt.Println("     OSTO CLI Authentication")
	fmt.Println("===================================")
	fmt.Println("Type 'help' to see available commands.")

	for {

		line, err := c.rl.Readline()
		if err != nil {
			break
		}

		switch line {

		case "register":
			c.authController.Register()

		case "login":
			c.authController.Login()

		case "whoami":
			c.authController.WhoAmI()

		case "enable-2fa":
			c.authController.Enable2FA()

		case "disable-2fa":
			c.authController.Disable2FA()

		case "logout":
			c.authController.Logout()

		case "help":
			c.printHelp()

		case "exit":
			fmt.Println("Goodbye!")
			return

		case "":
			continue

		default:
			fmt.Println("Unknown command. Type 'help'.")
		}
	}
}

func (c *CLI) printHelp() {

	fmt.Println()
	fmt.Println("Available Commands")
	fmt.Println("------------------")

	fmt.Println("register")
	fmt.Println("login")

	fmt.Println("whoami")
	fmt.Println("enable-2fa")
	fmt.Println("disable-2fa")
	fmt.Println("logout")

	fmt.Println("help")
	fmt.Println("exit")
}