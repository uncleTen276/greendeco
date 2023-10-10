package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
)

type MediaRoutes struct {
	app fiber.Router
}

func NewMediaRouter(app fiber.Router) *MediaRoutes {
	return &MediaRoutes{app: app.Group("/media")}
}

func (r *MediaRoutes) RegisterRoutes() {
	r.publicRouter()
}

func (r *MediaRoutes) publicRouter() {
	r.app.Post("/upload", controller.PostMedia)
}
