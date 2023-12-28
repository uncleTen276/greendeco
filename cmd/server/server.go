package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/sekke276/greendeco.git/docs"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/pkg/routes"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
	"github.com/sekke276/greendeco.git/platform/storage"
)

func Serve() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("error")
	}

	if err := database.ConnectDB(); err != nil {
		log.Panic(err)
	}

	if err := database.GetDB().Migrate(); err != nil {
		log.Panic(err)
	}

	if err := storage.ConnectStorage(); err != nil {
		log.Panic(err)
	}

	validators.AddProductQueryDecoder()
	validators.AddOrderQueryDecoder()

	app := fiber.New()
	app.Use(logger.New())

	corsApp := cors.ConfigDefault
	corsApp.AllowCredentials = true
	corsApp.AllowHeaders = "Origin, X-Requested-With, Content-Type, Accept, Authorization"
	app.Use(cors.New(corsApp))

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hellooo")
	})
	routes.SwaggerRoute(app)
	api := app.Group("/api/v1")
	routes.UserRoutes(api)
	routes.AuthRoutes(api)
	routes.NewMediaRouter(api).RegisterRoutes()
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
	routes.NewPaymentRouter(api).RegisterRoutes()
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("not response")
	}
}
