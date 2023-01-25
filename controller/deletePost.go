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

func DeletePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	postID, _ := strconv.Atoi(c.Params("id"))

	var postFormat models.Post
	var creatorID int64
	var userID int

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

	//Checks if user is indeed the post creator
	if err := connect.DB.Table("posts").Select("user_id").Where("id = ?", postID).Find(&creatorID); err != nil {
		fmt.Println("Error retrieving post's creator: delete post")
	}

	fmt.Println("issuer", claims.Issuer)
	fmt.Println("issuer", claims.Issuer)
	userID, _ = strconv.Atoi(claims.Issuer)

	if creatorID != int64(userID) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You are not the creator of this post.",
		})
	}

	//To prevent the foreign key constraint - cannot delete parent row error, we need to delete the posts' comments first
	var commentFormat models.Comment
	connect.DB.Where("post_id = ?", postID).Delete(&commentFormat)

	//Now we delete from posts
	connect.DB.Delete(&postFormat, postID)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Post deleted.",
	})
}
