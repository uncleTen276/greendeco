package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

func CategoryRouter(app fiber.Router) {
	category := app.Group("/category")
	publicCategoryRouter(category)
	privateCategoryRouter(category)
}

func publicCategoryRouter(app fiber.Router) {
	app.Get("/", controller.GetAllCategory)
}

func privateCategoryRouter(app fiber.Router) {
	app.Use(middlewares.JWTProtected())
	app.Use(middlewares.AdminProtected)
	app.Post("/", controller.CreateCategories)
	app.Put("/:id", controller.UpdateCategories)
	app.Delete("/:id", controller.DeleteCategories)
}
