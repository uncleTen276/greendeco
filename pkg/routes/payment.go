package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
)

type paymentRoutes struct {
	app fiber.Router
}

func NewPaymentRouter(app fiber.Router) *paymentRoutes {
	return &paymentRoutes{
		app: app.Group("/payment"),
	}
}

func (r *paymentRoutes) RegisterRoutes() {
	r.publicOrderRoutes()
	r.privateOrderRoutes()
}

func (r *paymentRoutes) publicOrderRoutes() {
	r.app.Get("/vnpay_return", controller.VnPay_Return)
	r.app.Post("/paypal_return", controller.PayPalReturn)
}

func (r *paymentRoutes) privateOrderRoutes() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Post("/vnpay_create", controller.CreateVnPayPayment)
	r.app.Post("/paypal_create", controller.CreatePayPalPayment)
}
