package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	// Create a rate limiter middleware
	// limiterConfig := limiter.Config{
	// 	Max:        5,               // Maximum requests allowed per window
	// 	Expiration: 1 * time.Minute, // Window duration
	// }
	// limiterMiddleware := limiter.New(limiterConfig)

	// Apply rate limiting middleware to specific paths
	app.Get("/", limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        1,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"message": fiber.ErrTooManyRequests.Message,
				"success": false,
				"data":    nil,
			})
		},
	}), getResource)
	app.Get("/test", limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        5,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"message": fiber.ErrTooManyRequests.Message,
				"success": false,
				"data":    nil,
			})
		},
	}), getSpesificResource)

	app.Listen(":3000")
}

func getResource(c *fiber.Ctx) error {
	return c.SendString("Accessing GET /")
}

func getSpesificResource(c *fiber.Ctx) error {
	return c.SendString("Accessing GET /get")
}

// func rateLimitGlobalConfig() limiter.Config {
// 	return limiter.Config{
// 		Expiration: 10 * time.Second,
// 		Max:        1,
// 		LimitReached: func(c *fiber.Ctx) error {
// 			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
// 				"message": fiber.ErrTooManyRequests.Message,
// 				"success": false,
// 				"data":    nil,
// 			})
// 		},
// 	}
// }
