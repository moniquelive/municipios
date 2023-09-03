package main

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"github.com/moniquelive/municipios/router"
)

func main() {
	engine := html.New("./views", ".gohtml")
	engine.Reload(true)
	engine.Debug(true)
	engine.AddFuncMap(sprig.HtmlFuncMap())
	fiberConfig := fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
	}
	app := fiber.New(fiberConfig)
	app.Use(favicon.New()).
		Use(logger.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen("localhost:9090"))
}
