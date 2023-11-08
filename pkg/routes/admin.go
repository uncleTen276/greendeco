package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type AdminRouters struct {
	app fiber.Router
}

func NewAdminRouter(app fiber.Router) *AdminRouters {
	return &AdminRouters{app: app.Group("/admin")}
}

func (r *AdminRouters) RegisterRoutes() {
	r.publisAdminRoute()
	r.privateAdminRoute()
}

func (r *AdminRouters) publisAdminRoute() {
	r.app.Get("/login", controller.LoginForAdmin)
}

func (r *AdminRouters) privateAdminRoute() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Use(middlewares.AdminProtected)
	r.app.Get("/customers", controller.GetAllCustomers)
}
