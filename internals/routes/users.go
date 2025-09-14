package routes

import "github.com/gofiber/fiber/v2"

func Users(r fiber.Router) {
	users := r.Group("/users")

	users.Post("/", nil)
	users.Get("/", nil)

	users.Get("/:id", nil)
	users.Put("/:id", nil)
	users.Delete("/:id", nil)
}
