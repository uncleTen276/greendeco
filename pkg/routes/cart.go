package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type CartRouters struct {
	app fiber.Router
}

func NewCartRouter(app fiber.Router) *CartRouters {
	return &CartRouters{
		app: app.Group("/cart"),
	}
}

func (r *CartRouters) RegisterRoutes() {
	r.privateProductRouter()
}

func (r *CartRouters) privateProductRouter() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Get("/", controller.GetCartByOnwerId)
	r.app.Get("/:id", controller.GetCartById)
	r.app.Get("/product/:id", controller.GetCartProductById)
	r.app.Get("/:id/product", controller.GetCartProductsByCartId)
	r.app.Post("/", controller.CreateCart)
	r.app.Post("/product/", controller.CreateCartProduct)
	r.app.Put("/product/:id", controller.UpdateCartProduct)
	r.app.Delete("/:id/clear", controller.DeleteCartItemByCartId)
	r.app.Delete("/product/:id", controller.DeleteCartItemById)
}
