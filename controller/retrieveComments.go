package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func RetrieveComments(c *fiber.Ctx) error {
	var data []models.Comment
	postID := c.Params("id")

	if err := connect.DB.Find(&data, "post_id = ?", postID); err != nil {
		fmt.Println("Error retrieving comments for post")
	}

	c.Status(200)
	return c.JSON(data)

}
