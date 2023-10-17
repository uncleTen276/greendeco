package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type ReviewRouters struct {
	app fiber.Router
}

func NewReviewRoutes(app fiber.Router) *ReviewRouters {
	return &ReviewRouters{app: app.Group("/review")}
}

func (r *ReviewRouters) RegisterRoutes() {
	r.publicReviewRouter()
	r.privateReviewRouter()
}

func (r *ReviewRouters) publicReviewRouter() {
	r.app.Get("/:id", controller.GetReviewById)
	r.app.Get("/product/:id", controller.GetReviewByProductId)
}

func (r *ReviewRouters) privateReviewRouter() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Post("/", controller.CreateReview)
}
