package routes

import (
    "go-fiber-mvc/controllers"
    "go-fiber-mvc/services"
	"go-fiber-mvc/middlewares"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
    userService := services.NewUserService(db)
    userController := controllers.NewUserController(userService)
    authController := controllers.NewAuthController(db)

    // Public routes
	app.Post("/register", authController.Register)
    app.Post("/login", authController.Login)
    app.Get("/users", userController.GetUsers)
    app.Get("/users/:id", userController.GetUser)

    // Protected routes
    app.Post("/users", middlewares.AuthMiddleware, userController.CreateUser)
    app.Put("/users/:id", middlewares.AuthMiddleware, userController.UpdateUser)
    app.Delete("/users/:id", middlewares.AuthMiddleware, userController.DeleteUser)
}
