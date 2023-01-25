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

func UpdatePost(c *fiber.Ctx) error {
	//jwt needed to ensure that the user is the creator of the edited post
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil
	})

	//error handling
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	//Access issuer through claims.Issuer
	claims := token.Claims.(*jwt.StandardClaims)

	var updatedPost map[string]interface{}

	if err := c.BodyParser(&updatedPost); err != nil {
		fmt.Println("Create Post: Unable to parse body")
	}

	var postNew models.Post

	//convert the postid received to int
	postIDint, _ := strconv.Atoi(updatedPost["postid"].(string))

	//retrieves the post with given id
	if err := connect.DB.First(&postNew, postIDint); err != nil {
		fmt.Println("Error retrieving post for update")
	}

	userIDint, _ := strconv.Atoi(claims.Issuer)

	//Checks if the posts creator id matches the token user id.
	if postNew.UserID != uint(userIDint) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You may not update this post as you are not its creator.",
		})
	}
	postNew.Title = updatedPost["title"].(string)
	postNew.Desc = updatedPost["desc"].(string)
	postNew.Content = updatedPost["content"].(string)

	//updates the post
	if err := connect.DB.Save(&postNew); err != nil {
		fmt.Println("Error saving updated post")
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Post updated successfully!"})
}
