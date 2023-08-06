package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/web"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hellooo")
	})
	web.Routes(app)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("not response")
	}
}
