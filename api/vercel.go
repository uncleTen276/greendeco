package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sekke276/greendeco.git/web"
)

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
