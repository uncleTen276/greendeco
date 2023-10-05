package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

func ProductRouter(app fiber.Router) {
	category := app.Group("/category")
	privateProductRouter(category)
	publicProductRouter(category)
}

func publicProductRouter(app fiber.Router) {
}

func privateProductRouter(app fiber.Router) {
	app.Use(middlewares.JWTProtected())
	app.Post("/", controller.CreateCategories)
	app.Post("/:id/update", controller.UpdateCategories)
	app.Delete("/:id/delete", controller.DeleteCategories)
}
