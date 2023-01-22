package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func GetCommunity(c *fiber.Ctx) error {
	type CommReq struct {
		Link string `json:"commlink"`
	}

	var reqComm CommReq

	if err := c.BodyParser(&reqComm); err != nil {
		fmt.Println("Request for Community: Unable to parse body")
	}

	var commID int

	//retrieving id of the community
	if err := connect.DB.Table("communities").Select("id").Where("link = ?", reqComm.Link).Find(&commID); err != nil {
		fmt.Println("Error retrieving community ID")
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

	return (c.JSON(postList))
}
