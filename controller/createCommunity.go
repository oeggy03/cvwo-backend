package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func CreateCommunity(c *fiber.Ctx) error {
	var newCommunity models.Community
	var checkCommunity models.Community

	if err := c.BodyParser(&newCommunity); err != nil {
		fmt.Println("Create Community: Unable to parse body")
	}

	//checks if community already exists
	connect.DB.Where("link = ?", newCommunity.Link).First(&checkCommunity)

	if checkCommunity.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Community" + newCommunity.Name + "already exist",
		})
	}

	if err := connect.DB.Create(&newCommunity).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create Community: Invalid payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Community Created!",
	})
}
