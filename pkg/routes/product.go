package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type ProductRouters struct {
	app fiber.Router
}

func NewProductRouter(app fiber.Router) *ProductRouters {
	return &ProductRouters{app: app.Group("/product")}
}

func (r *ProductRouters) RegisterRoutes() {
	r.publicProductRouter()
	r.privateProductRouter()
}

func (r *ProductRouters) publicProductRouter() {
	// app.Get("/",)
}

func (r *ProductRouters) privateProductRouter() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Post("/", controller.CreateProduct)
	r.app.Post("/:id/update", controller.UpdateProduct)
	r.app.Delete("/:id/delete", controller.DeleteProduct)
}
