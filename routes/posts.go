package routes

import "github.com/gofiber/fiber/v2"

func Posts(r fiber.Router) {

	posts := r.Group("/posts")
	posts.Post("/", nil)
	posts.Get("/", nil)

	posts.Get("/:id", nil)
	posts.Delete("/:id", nil)

	//TODO : update posts
}
