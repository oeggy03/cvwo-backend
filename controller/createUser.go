package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	var checkUser models.User

	if err := c.BodyParser(&user); err != nil {
		fmt.Println("Create User: Unable to parse body")
	}

	connect.DB.Where("username = ?", user.Username).First(&user)

	if checkUser.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": user.Username + "already exist",
		})
	}

	if err := connect.DB.Create(&user).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create User: Invalid payload",
		})
	}

	c.Status(200)
	return c.JSON("User " + user.Username + " created successfully.")
}
