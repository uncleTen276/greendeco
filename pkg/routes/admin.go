package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

func AdminRoute(app fiber.Router) {
	admin := app.Group("/admin")
	privateAdminRoute(admin)
}

func privateAdminRoute(app fiber.Router) {
	app.Use(middlewares.JWTProtected())
	app.Get("/customers", controller.GetAllCustomers)
}
