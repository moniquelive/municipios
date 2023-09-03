// Package router is a single function package that sets up default routes
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/moniquelive/municipios/handler"
)

// SetupRoutes creates default routes
func SetupRoutes(app *fiber.App) {
	app.Static("/css", "./web/css").
		Static("/js", "./web/js").
		Static("/images", "./web/images")

	app.Get("/status", monitor.New())
	app.Get("/", handler.Index)
	app.Get("/estado", handler.Estado)
}
