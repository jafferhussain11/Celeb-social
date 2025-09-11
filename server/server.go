package server

import "github.com/gofiber/fiber/v2"

func setup(app *fiber.App) {
	app = fiber.New(fiber.Config{
		ErrorHandler: nil,
		BodyLimit:    1024 * 1024 * 16,
	})
}
