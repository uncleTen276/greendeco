package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/pkg/routes"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
	"github.com/sekke276/greendeco.git/platform/storage"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// @title Greendeco API
// @version 1.0
// @description greendeco
// @contact.name Nguyen Tri
// @contact.email tringuyen2762001@gmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath  /api/v1
func handler() http.HandlerFunc {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("error")
	}

	if err := database.ConnectDB(); err != nil {
		log.Panic(err)
	}

	if err := storage.ConnectStorage(); err != nil {
		log.Panic(err)
	}

	validators.AddProductQueryDecoder()
	validators.AddOrderQueryDecoder()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.SwaggerRoute(app)
	api := app.Group("/api/v1")
	routes.AuthRoutes(api)
	routes.UserRoutes(api)
	routes.NewMediaRouter(api).RegisterRoutes()
	routes.NewProductRouter(api).RegisterRoutes()
	routes.NewAdminRouter(api).RegisterRoutes()
	routes.CategoryRouter(api)

	routes.NewReviewRoutes(api).RegisterRoutes()
	routes.NewProductRouter(api).RegisterRoutes()
	routes.NewVariantRouter(api).RegisterRoute()
	routes.NewCartRouter(api).RegisterRoutes()
	routes.NewColorRouter(api).RegisterRoutes()
	routes.NewCouponRouter(api).RegisterRoutes()
	routes.NewOrderRouter(api).RegisterRoutes()
	routes.NewNotificationRouter(api).RegisterRoutes()

	return adaptor.FiberApp(app)
}
