// main initializes the application dependencies,
// establishes the database connection,
// and starts the interactive CLI.
package main

import (
	"log"

	"github.com/Aditya7880900936/osto-cli-login/internal/cli"
	"github.com/Aditya7880900936/osto-cli-login/internal/controllers"
	"github.com/Aditya7880900936/osto-cli-login/internal/database"
	"github.com/Aditya7880900936/osto-cli-login/internal/repository"
	"github.com/Aditya7880900936/osto-cli-login/internal/services"
	"github.com/Aditya7880900936/osto-cli-login/internal/session"
)

func main() {

	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository()

	authService := services.NewAuthService(userRepository)

	sessionManager := session.NewSessionManager()

	authController := controllers.NewAuthController(
		authService,
		sessionManager,
	)

	app := cli.NewCLI(authController)

	app.Start()
}
