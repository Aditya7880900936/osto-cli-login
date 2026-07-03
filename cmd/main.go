package main

import (
	"log"

	"github.com/Aditya7880900936/osto-cli-login/internal/cli"
	"github.com/Aditya7880900936/osto-cli-login/internal/controllers"
	"github.com/Aditya7880900936/osto-cli-login/internal/database"
	"github.com/Aditya7880900936/osto-cli-login/internal/repository"
	"github.com/Aditya7880900936/osto-cli-login/internal/services"
)

func main() {

	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository()

	authService := services.NewAuthService(userRepository)

	authController := controllers.NewAuthController(authService)

	application := cli.NewCLI(authController)

	application.Start()
}