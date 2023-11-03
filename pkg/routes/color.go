package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type ColorRoutes struct {
	app fiber.Router
}

func NewColorRouter(app fiber.Router) *ColorRoutes {
	return &ColorRoutes{
		app: app.Group("/color"),
	}
}

func (r *ColorRoutes) RegisterRoutes() {
	r.publishColorRouter()
	r.privateColorRouter()
}

func (r *ColorRoutes) publishColorRouter() {
	r.app.Get("/", controller.GetColors)
	r.app.Get("/:id", controller.GetColorById)
}

func (r *ColorRoutes) privateColorRouter() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Use(middlewares.AdminProtected)
	r.app.Post("/", controller.CreateColor)
	r.app.Put("/:id", controller.UpdateColorById)
}
