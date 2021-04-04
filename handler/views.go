package handler

import "github.com/gofiber/fiber"

// GetViews for views
func GetViews(c *fiber.Ctx) {
	_ = c.Render("index", fiber.Map{
		"Title": "Hello, Golang Client!",
	})
}
