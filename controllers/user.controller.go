package controllers

import (
	"belajar-golang-fiber/database"
	"belajar-golang-fiber/model/entity"

	"github.com/gofiber/fiber/v2"
)

// GetUsers returns a list of users.
func GetUsers(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return result.Error

	}
	return ctx.JSON(&fiber.Map{
		"status": "success",
		"data":   users,
	})
}

// GetUser returns the user with the given ID.
func GetUser(ctx *fiber.Ctx) error {
	var user entity.User
	result := database.DB.First(&user, ctx.Params("id"))
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSON(user)
}

// CreateUser creates a new user.
func CreateUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Create User")
}

// UpdateUser updates the user with the given ID.
func UpdateUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User" + ctx.Params("id"))
}

// DeleteUser deletes the user with the given ID.
func DeleteUser(ctx *fiber.Ctx) error {
	var user entity.User
	res := database.DB.Delete(&user, ctx.Params("id"))
	if res.Error != nil {
		return res.Error
	}
	return ctx.JSON(user)
}
