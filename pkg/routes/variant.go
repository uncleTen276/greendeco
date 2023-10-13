package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type VariantRoutes struct {
	app fiber.Router
}

func NewVariantRouter(app fiber.Router) *VariantRoutes {
	return &VariantRoutes{app: app.Group("/variant")}
}

func (r *VariantRoutes) RegisterRoute() {
	r.publicProductRouter()
	r.privateProductRouter()
}

func (r *VariantRoutes) publicProductRouter() {
	r.app.Get("/product/:id", controller.GetVariantsByProductId)
}

func (r *VariantRoutes) privateProductRouter() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Post("/", controller.CreateVariant)
}
