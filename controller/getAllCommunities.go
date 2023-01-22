package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func GetAllCommunities(c *fiber.Ctx) error {
	var CommList []models.Community

	if err := connect.DB.Find(&CommList); err != nil {
		c.Status(400)
		fmt.Println(err)
	}

	c.Status(200)
	return c.JSON(CommList)
}
