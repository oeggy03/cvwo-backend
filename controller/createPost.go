package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		fmt.Println("Create Post: Unable to parse body")
	}

	if err := connect.DB.Create(&post).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create Post: Invalid payload",
		})
	}

	c.Status(200)
	return c.JSON("Post " + post.Title + " created successfully.")
}
