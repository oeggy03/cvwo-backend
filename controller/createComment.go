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

func CreateComment(c *fiber.Ctx) error {
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

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Create Post: Unable to parse body")
	}

	intPostID, _ := strconv.Atoi(data["postid"])
	intUserID, _ := strconv.Atoi(claims.Issuer)

	//To assign the creator username to comment
	var creator models.User
	connect.DB.Where("id = ?", claims.Issuer).First(&creator)

	comment := models.Comment{
		Content: data["content"],
		PostID:  intPostID,
		UserID:  intUserID,
		Creator: creator.Username,
	}

	if err := connect.DB.Create(&comment).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create Comment: Invalid payload",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
