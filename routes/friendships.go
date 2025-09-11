package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Friendships(r fiber.Router) {

	friendships := r.Group("/friendships")

	friendships.Post("/", nil)
	friendships.Get("/", nil)

	friendships.Delete("/:id", nil)
}
