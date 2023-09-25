package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/sekke276/greendeco.git/docs"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/pkg/routes"
	"github.com/sekke276/greendeco.git/platform/database"
	"github.com/sekke276/greendeco.git/platform/storage"
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

	if err := storage.ConnectStorage(); err != nil {
		log.Panic(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hellooo")
	})
	routes.SwaggerRoute(app)
	api := app.Group("/api/v1")
	routes.UserRoutes(api)
	routes.AuthRoutes(api)
	routes.MediaRoutes(api)
	routes.AdminRoute(api)
	web.Routes(app)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("not response")
	}
}
