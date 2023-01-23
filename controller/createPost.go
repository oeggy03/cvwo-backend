package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func CreatePost(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Create Post: Unable to parse body")
	}

	fmt.Println(data)

	intID, _ := strconv.Atoi(data["userid"].(string))
	// intComm, _ := strconv.Atoi(data["communityid"].(string))

	post := models.Post{
		Title:       data["title"].(string),
		Desc:        data["desc"].(string),
		Content:     data["content"].(string),
		UserID:      uint(intID),
		CommunityID: uint(data["communityid"].(float64)),
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
