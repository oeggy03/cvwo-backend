package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/oeggy03/cvwo-backend/connect"
	"github.com/oeggy03/cvwo-backend/routes"
)

func main() {
	connect.Connect()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env files - port")
	} else {
		log.Println("Connected successfully")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.Setup(app)
	app.Listen(":" + port)

}
