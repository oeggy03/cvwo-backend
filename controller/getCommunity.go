package controller

import "github.com/gofiber/fiber/v2"

func GetCommunity(c *fiber.Ctx) error {
	return (c.JSON(fiber.Map{
		"hello": "hello",
	}))
}
