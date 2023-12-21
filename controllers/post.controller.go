package controllers

import "github.com/gofiber/fiber/v2"

// GetPosts returns all posts
func GetPosts(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Posts")
}

// GetPost returns a specific post
func GetPost(ctx *fiber.Ctx) error {
	return ctx.Send([]byte(ctx.Params("id")))
}

// CreatePost creates a new post
func CreatePost(ctx *fiber.Ctx) error {
	return ctx.SendString("Create Post")
}

// UpdatePost updates an existing post
func UpdatePost(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Post " + ctx.Params("id"))
}

// DeletePost deletes a post
func DeletePost(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete Post" + ctx.Params("id"))
}
