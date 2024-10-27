package controllers

import (
    "myapp/middlewares"
    "myapp/models"
    "myapp/services"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type AuthController struct {
    DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
    return &AuthController{DB: db}
}

// Login function to authenticate users and return a JWT token
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
    type LoginRequest struct {
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required"`
    }

    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    // Validate the request data
    if err := services.ValidateStruct(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"errors": err})
    }

    // Find user by email
    var user models.User
    if err := ctrl.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
    }

    // Check password (assuming passwords are hashed)
    if user.Password != req.Password { // Replace with hash check
        return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
    }

    // Generate JWT token
    token, err := middlewares.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}

// Register function to create a new user
func (ctrl *AuthController) Register(c *fiber.Ctx) error {
    type RegisterRequest struct {
        Name     string `json:"name" validate:"required"`
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required,min=6"`
    }

    var req RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    // Validate the request data
    if err := services.ValidateStruct(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"errors": err})
    }

    // Check if user already exists
    var existingUser models.User
    if err := ctrl.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        return c.Status(400).JSON(fiber.Map{"error": "Email already in use"})
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
    }

    // Create user
    user := models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: string(hashedPassword),
    }

    if err := ctrl.DB.Create(&user).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
    }

    // Generate JWT token for the new user
    token, err := middlewares.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{"message": "User registered successfully", "token": token})
}
