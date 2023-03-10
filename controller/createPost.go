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

func CreatePost(c *fiber.Ctx) error {
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

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Create Post: Unable to parse body")
	}

	intID, _ := strconv.Atoi(claims.Issuer)

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
	return c.JSON(post.ID)
}
