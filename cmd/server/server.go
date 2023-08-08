package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/platform/database"
	"github.com/sekke276/greendeco.git/web"
)

func Serve() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("error")
	}

	if err := database.ConnectDB(cfg); err != nil {
		log.Panic(err)
	}

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hellooo")
	})
	web.Routes(app)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("not response")
	}
}
