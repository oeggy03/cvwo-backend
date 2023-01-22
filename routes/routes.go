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
	app.Get("/api/GetCommunities/:id", controller.GetCommunity)
	app.Post("/api/CreateUser", controller.CreateUser)
	// //when we go to this link, register is called
	// app.Post("/api/register", controller.Register)
	// app.Post("/api/login", controller.Login)
	// //If user is not authenticated, they cannot access any route below this middleware
	// app.Use(middleware.IsAuthenticated)
	// app.Post("/api/post", controller.CreatePost)
	// app.Get("/api/allposts", controller.AllPost)
	// app.Get("/api/allposts/:id", controller.DetailedPost)
	// app.Put("/api/updatepost/:id", controller.UpdatePost)
	// app.Get("/api/uniquepost", controller.UniquePost)
	// app.Delete("/api/deletepost/:id", controller.DeletePost)
	// app.Post("/api/upload-image", controller.Upload)
	// app.Static("/api/uploads", "./uploads")
}
