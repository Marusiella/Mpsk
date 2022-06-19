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

	// v1.Get("/ttt", func(c *fiber.Ctx) error {
	// 	token := c.Get("JWT")
	// 	user := &Claims{}
	// 	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
	// 		return []byte(SECRET_JWT), nil
	// 	})
	// 	if err != nil {
	// 		return c.Status(401).SendString("Unauthorized")
	// 	}
	// 	var grupa []Group
	// 	db.Preload("Users").Preload("Tasks").Find(&grupa)
	// 	return c.JSON(grupa)
	// })
	log.Fatal(app.Listen(":3000"))
}
