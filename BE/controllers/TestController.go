package controllers

import "github.com/gofiber/fiber/v2"

func TestController(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"hello": "world",
	})

}
