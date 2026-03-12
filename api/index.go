package api

import (
	"embed"
	"net/http"
	"portofolio/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

//go:embed all:views
var viewsFS embed.FS

//go:embed all:public
var publicFS embed.FS

var app *fiber.App

func init() {
	// Initialize template engine using embedded files
	engine := html.NewFileSystem(http.FS(viewsFS), ".html")

	// Create Fiber app
	app = fiber.New(fiber.Config{
		Views: engine,
	})

	// Static file handler using embedded files
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(publicFS),
		PathPrefix: "public",
		Browse:     false,
	}))

	// Setup routes
	routes.SetupRoutes(app)
}

// Handler is the exported function Vercel will call
func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}
