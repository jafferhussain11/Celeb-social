package server

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

var app *fiber.App

// TODO : maybe use sync.once() for singleton
func New() *fiber.App {
	return app
}

func Setup() {
	app = fiber.New(fiber.Config{
		ErrorHandler: nil,
		BodyLimit:    1024 * 1024 * 16,
	})

	defer app.Use(notFoundHandler)
	defer app.Use(recover2.New())

	middlewares(app)
	addRoutes(app)
}
