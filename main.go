package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Details struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsPresent bool   `json:"present"`
}

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

	app.Post("/v2/details", func(c *fiber.Ctx) error {
		detail := new(Details)
		err := c.BodyParser(detail)
		if err != nil {
			return err
		}
		log.Println(detail.Name)
		return c.SendString("details found")
	})

	app.Listen(":3000")
}
