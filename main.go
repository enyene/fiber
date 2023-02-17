package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("welcome")
	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	// ...

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("hello " + c.Params("value"))
	})

	app.Get("/:name?", func(c *fiber.Ctx) error {
		return c.SendString("hello" + c.Params("name"))
	})
	app.Listen(":3000")
}
