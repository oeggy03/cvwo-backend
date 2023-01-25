package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/cvwo-backend/controller"
	// "github.com/oeggy03/blog-backend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/CreateCommunity", controller.CreateCommunity)
	app.Get("/api/GetCommunities", controller.GetAllCommunities)
	app.Post("/api/CreatePost", controller.CreatePost)
	app.Get("/api/GetCommunity/:link", controller.GetCommunity)
	app.Get("/api/GetCommDetails/:link", controller.GetCommDetails)
	app.Post("/api/CreateUser", controller.CreateUser)
	app.Post("/api/SignIn", controller.SignIn)
	app.Get("/api/GetUser", controller.GetUser)
	app.Get("/api/GetUserPosts", controller.GetUserPosts)
	app.Post("/api/LogOut", controller.LogOut)
	app.Get("/api/RetrievePost/:id", controller.RetrievePost)
	app.Delete("/api/DeletePost/:id", controller.DeletePost)
	app.Post("/api/CreateComment", controller.CreateComment)
	app.Get("/api/RetrieveComments/:id", controller.RetrieveComments)
	app.Put("/api/UpdatePost", controller.UpdatePost)
	app.Delete("/api/DeleteComment/:id", controller.DeleteComment)
	app.Put("/api/UpdateComment", controller.EditComment)
}
