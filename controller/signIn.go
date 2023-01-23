package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/connect"
	util "github.com/oeggy03/cvwo-backend/jwt"
	"github.com/oeggy03/cvwo-backend/models"
)

func SignIn(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		fmt.Println("Unable to parse body")
	}

	var user models.User

	//check if username exists in our database
	connect.DB.Where("username=?", data["username"]).First(&user)
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Sorry, user not found.",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password.",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa((int(user.ID))))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Login Successful! You may close the popup.",
	})

}

type Claims struct {
	jwt.StandardClaims
}
