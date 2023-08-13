package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/sekke276/greendeco.git/docs"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/pkg/routes"
	"github.com/sekke276/greendeco.git/platform/database"
	"github.com/sekke276/greendeco.git/web"
)

func Serve() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("error")
	}

	if err := database.ConnectDB(); err != nil {
		log.Panic(err)
	}
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hellooo")
	})
	routes.SwaggerRoute(app)
	web.Routes(app)
	api := app.Group("/api/v1")
	routes.AuthRoutes(api)
	routes.NotFoundRoute(api)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("not response")
	}
}
