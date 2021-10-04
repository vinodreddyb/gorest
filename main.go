package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"restexample/handlers"
	"restexample/repository"
)

func main() {
	_, err := repository.Connect()

	if err != nil {
		log.Fatal(err)
	}
	// Create a new Fiber instance
	app := fiber.New()
	// Use logger
	app.Use(logger.New())
	userGroup := app.Group("/user")
	userGroup.Post("/", handlers.AddNewUser)
	userGroup.Get("/", handlers.GetAllUsers)
	err = app.Listen(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
