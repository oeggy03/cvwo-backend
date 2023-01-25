package controller

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func CreateUser(c *fiber.Ctx) error {
	var data map[string]interface{}
	var checkUser models.User
	var checkEmail models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Create User: Unable to parse body")
	}

	connect.DB.Where("username = ?", data["username"].(string)).First(&checkUser)

	if len(data["username"].(string)) == 0 || len(data["password"].(string)) == 0 || len(data["email"].(string)) == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Please do not leave any field empty.",
		})
	}
	if checkUser.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sorry, the username " + data["username"].(string) + " is taken.",
		})
	}

	connect.DB.Where("email = ?", data["email"].(string)).First(&checkEmail)

	if checkEmail.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sorry, " + data["email"].(string) + " has already been used.",
		})
	}

	user := models.User{
		Username: data["username"].(string),
		Email:    strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))

	if err := connect.DB.Create(&user).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create User: Invalid payload.",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Success! Please close this popup and sign in.",
	})
}
