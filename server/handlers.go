package server

import "github.com/gofiber/fiber/v2"

//type ErrorHandler = func(*Ctx, error) error

func errorHandler(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	msg := fiber.Map{
		"status":  "error",
		"message": err.Error(),
	}
	return c.Status(fiber.StatusInternalServerError).JSON(msg)
}

var notFoundHandler = func(c *fiber.Ctx) error {
	msg := fiber.Map{
		"status":  "error",
		"message": "Requested resource at Route " + c.OriginalURL() + " Not Found",
	}
	return c.Status(fiber.StatusNotFound).JSON(msg)
}
