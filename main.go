package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "welcome to fiber",
		})
	})

	app.Get("/api/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Hello: " + name,
		})
	})

	app.Listen(":3000")
}
