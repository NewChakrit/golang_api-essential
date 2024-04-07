package views

import "github.com/gofiber/fiber/v2"

func TestHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":  "Hello, World!",
		"Title2": "My name's New :)",
	})
}
