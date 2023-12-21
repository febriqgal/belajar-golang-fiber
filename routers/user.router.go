package routers

import (
	"belajar-golang-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterUser(c *fiber.App, api fiber.Router) {
	api.Get("/users", controllers.GetUsers)
	api.Get("/user/:id", controllers.GetUser)
	api.Post("/user", controllers.CreateUser)
	api.Put("/user/:id", controllers.UpdateUser)
	api.Delete("/user/:id", controllers.DeleteUser)
}
