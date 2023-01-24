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

func RetrieveComments(c *fiber.Ctx) error {
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

	var data []models.Comment
	postID := c.Params("id")

	if err := connect.DB.Order("id desc").Find(&data, "post_id = ?", postID); err != nil {
		fmt.Println("Error retrieving comments for post")
	}

	intUserid, _ := strconv.Atoi(claims.Issuer)

	c.Status(200)
	return c.JSON(fiber.Map{
		"userid":   intUserid,
		"comments": data,
	})

}
