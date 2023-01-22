package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func GetCommunity(c *fiber.Ctx) error {
	link := c.Params("link")

	var commID int
	var comm models.Community

	//retrieving id of the community
	if err := connect.DB.Table("communities").Select("id").Where("link = ?", link).Find(&commID); err != nil {
		fmt.Println("Error retrieving community ID")
	}

	if err := connect.DB.First(&comm, "id = ?", commID); err != nil {
		fmt.Println("Error retrieving community")
	}

	if commID == 0 {
		c.Status(400)
		return (c.JSON(fiber.Map{
			"message": "Community does not exist.",
		}))
	}

	//retrieving all posts from that community
	var postList []models.Post

	if err := connect.DB.Find(&postList, "community_id = ?", commID); err != nil {
		fmt.Println("Error retrieving posts from community")
	}

	return (c.JSON(fiber.Map{
		"community": comm,
		"posts":     postList}))
}
