package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:dist
var NextFX embed.FS

func Routes(app fiber.Router) {
	dist, err := fs.Sub(NextFX, "dist")
	if err != nil {
		log.Fatal(err)
	}

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(dist),
		Browse:       true,
		NotFoundFile: "404.html",
	}))
}
