package routers

import (
	"belajar-golang-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterPost(c *fiber.App, api fiber.Router) {
	api.Get("/posts", controllers.GetPosts)
	api.Get("/post/:id", controllers.GetPost)
	api.Post("/post", controllers.CreatePost)
	api.Put("/post/:id", controllers.UpdatePost)
	api.Delete("/post/:id", controllers.DeletePost)

}
