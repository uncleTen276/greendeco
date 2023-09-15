package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
)

func MediaRoutes(app fiber.Router) {
	media := app.Group("/media")
	media.Post("/upload", controller.PostMedia)
}
