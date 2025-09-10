package main

import (
	"fmt"

	"github.com/jafferhussain11/celeb-social/cmd/app"
	"github.com/joho/godotenv"
)

func main() {

	//load env variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	app.Setup()

}
