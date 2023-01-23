package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func GetCommDetails(c *fiber.Ctx) error {
	var comm models.Community
	link := c.Params("link")

	if err := connect.DB.First(&comm, "link = ?", link); err != nil {
		fmt.Println("Error retrieving community")
	}

	if comm.ID == 0 {
		c.Status(404)
		return c.JSON("Community not found.")
	}

	c.Status(200)
	return c.JSON(comm)

}
