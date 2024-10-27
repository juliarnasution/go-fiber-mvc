package main

import (
    "go-fiber-mvc/config"
    "go-fiber-mvc/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Initialize Database
    db := config.InitDatabase()

    // Set up routes with database injection
    routes.SetupRoutes(app, db)

    // Start server
    app.Listen(":3000")
}
