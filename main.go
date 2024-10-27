package main

import (
    "go-fiber-mvc/config"
    "go-fiber-mvc/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new HTML engine
	engine := html.New("./views", ".html")
	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

    // Initialize Database
    db := config.InitDatabase()

    // Set up routes with database injection
    routes.SetupRoutes(app, db)

    // Start server
    app.Listen(":3000")
}
