package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/sekke276/greendeco.git/docs"
)

func SwaggerRoute(a fiber.Router) {
	a.Get("/docs/*", swagger.HandlerDefault)
	a.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})
}
