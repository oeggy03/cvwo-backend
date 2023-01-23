package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
)

func RetrievePost(c *fiber.Ctx) error {
	var post models.Post
	postID := c.Params("id")
	var comm string
	var user string

	if err := connect.DB.First(&post, "id = ?", postID); err != nil {
		fmt.Println("Error retrieving post")
	}

	if post.ID == 0 {
		c.Status(404)
		return c.JSON("Post not found.")
	}

	fmt.Println(post)
	if err := connect.DB.Table("communities").Select("name").Where("id = ?", post.CommunityID).Find(&comm); err != nil {
		fmt.Println("Error retrieving community: retrieve post")
	}

	if err := connect.DB.Table("users").Select("username").Where("id = ?", post.UserID).Find(&user); err != nil {
		fmt.Println("Error retrieving community: retrieve post")
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user":      user,
		"post":      post,
		"community": comm,
	})
}
