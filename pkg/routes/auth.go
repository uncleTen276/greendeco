package routers

import "github.com/gofiber/fiber/v2"

func AuthRoutes(app fiber.Router) {
	_ = app.Group("/auth")
}
