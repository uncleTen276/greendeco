package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/controller"
)

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")
	publicAuthRouter(auth)
	privateAuthRouter(auth)
}

func publicAuthRouter(app fiber.Router) {
	app.Post("/register", controller.CreateUser)
	app.Post("/login", controller.Login)
	app.Post("/refresh-token", controller.RefreshToken)
	app.Post("/forgot-password", controller.ForgotPassword)
}

func privateAuthRouter(app fiber.Router) {
}