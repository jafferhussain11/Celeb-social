package app

import (
	"log"

	"github.com/jafferhussain11/celeb-social/internals/database"
	"github.com/jafferhussain11/celeb-social/server"
)

func Setup() {
	database.Config()

	server.Setup()
	app := server.New()

	if err := app.Listen(":8090"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
