package main

import (
	"filemaod/database"
	"filemaod/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// TODO: make passwords safer
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 512, // 512MB
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
		},
	))
	database.Connect()
	database.FirstUser()
	routes.Configure(app)

	log.Fatal(app.Listen(":3000"))
}
