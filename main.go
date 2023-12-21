package main

import (
	"belajar-golang-fiber/database"
	"belajar-golang-fiber/database/migration"
	"belajar-golang-fiber/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectToDatabase()
	migration.RunMigration()
	app := fiber.New()
	routers.Router(app)
	app.Listen(":3000")
}
