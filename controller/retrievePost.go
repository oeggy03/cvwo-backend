package controller

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/models"
	"github.com/oeggy03/cvwo-backend/util"
)

func RetrievePost(c *fiber.Ctx) error {
	var post models.Post
	postID := c.Params("id")
	var comm string
	var user string

	//retrieving the jwt so that we may verify if our user is the creator of the original post.
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil
	})

	//error handling
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		fmt.Println("hello", err.Error())
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

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
		"owner":     claims.Issuer == strconv.Itoa(int(post.UserID)),
	})
}
