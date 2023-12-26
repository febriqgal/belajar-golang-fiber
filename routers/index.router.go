package routers

import (
	"github.com/gofiber/fiber/v2"
)

func Router(ctx *fiber.App) {
	api := ctx.Group("/api") //* /api
	v1 := api.Group("/v1")   //* /api/v1
	RouterPost(ctx, v1)      //* /api/v1/post
	RouterUser(ctx, v1)      //* /api/v1/user
}
