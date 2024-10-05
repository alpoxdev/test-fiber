package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"test-fiber/config"
	"test-fiber/database"
	"test-fiber/routes"
)

func setupRoutes(app *fiber.App) {
	app.Static("/", "./public")
}

func main() {
	config.Init()
	database.Init()
	database.Migrate()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	routes.Init(app)

	// Cors middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Error handler middleware
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // -> 404 Not Found
	})

	log.Info("Server is running on port 8000")
	log.Fatal(app.Listen(":8000"))
}