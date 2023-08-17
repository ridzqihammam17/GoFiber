package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	app.Use(
		limiter.New(limiter.Config{
			Expiration: 10 * time.Second,
			Max:        1,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
					"message": fiber.ErrTooManyRequests.Message,
					"success": false,
					"data":    nil,
				})
			},
		}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
