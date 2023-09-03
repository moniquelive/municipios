package handler

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Teste de Modal",
	}, "layouts/main")
}

func Estado(c *fiber.Ctx) error {
	return c.Render("municipios/"+c.Query("uf"), fiber.Map{})
}
