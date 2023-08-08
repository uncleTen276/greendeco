package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sekke276/greendeco.git/web"
)

// @title Fiber Go API
// @version 1.0
// @description greendeco
// @contact.name Nguyen Tri
// @contact.email tringuyen2762001@gmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /api
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	web.Routes(app)
	// app.Get("/docs/*", swagger.HandlerDefault)
	// app.Get("/api/", func(c *fiber.Ctx) error {
	// 	return c.Redirect("/docs")
	// })

	return adaptor.FiberApp(app)
}
