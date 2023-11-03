package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type NotificationRoutes struct {
	app fiber.Router
}

func NewNotificationRouter(app fiber.Router) *NotificationRoutes {
	return &NotificationRoutes{app: app.Group("/notification")}
}

func (r *NotificationRoutes) RegisterRoutes() {
	r.publicRouter()
	r.privateRouter()
}

func (r *NotificationRoutes) publicRouter() {
}

func (r *NotificationRoutes) privateRouter() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Get("/", controller.GetNotificationByToken)
	r.app.Put("/:id/user", controller.UpdateReadNotification)
	r.app.Get("/:id", controller.GetNotificationById)
	r.app.Use(middlewares.AdminProtected)
	r.app.Post("/", controller.CreateNotification)
	r.app.Post("/send/", controller.SendNotification)
	r.app.Put("/:id", controller.UpdateNotification)
}
