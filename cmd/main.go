package main

import (
	"fmt"
	"log"

	"github.com/jafferhussain11/celeb-social/cmd/app"
	"github.com/jafferhussain11/celeb-social/cmd/app/database"
	"github.com/jafferhussain11/celeb-social/cmd/app/server"
	"github.com/jafferhussain11/celeb-social/internals/repository"
	"github.com/joho/godotenv"
)

func main() {

	//load env variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	db := database.Config()

	//repos
	repos := repository.NewRepositories(db)

	//fiber app setup
	server.Setup()
	app := server.New() //fiber app

	if err := app.Listen(":8090"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
