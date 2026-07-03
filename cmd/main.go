package main

import (
	"fmt"
	"log"

	"github.com/Aditya7880900936/osto-cli-login/internal/database"
)

func main() {

	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully.")
}