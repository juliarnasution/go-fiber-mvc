package controllers

import (
    "go-fiber-mvc/models"
    "go-fiber-mvc/services"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type UserController struct {
    UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{UserService: userService}
}

func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
    users, err := ctrl.UserService.GetAllUsers()
    if err != nil {
        return c.Status(500).SendString("Failed to retrieve users")
    }
    return c.JSON(users)
}

func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    user, err := ctrl.UserService.GetUserByID(uint(id))
    if err != nil {
        return c.Status(404).SendString("User not found")
    }
    return c.JSON(user)
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString("Invalid request payload")
    }

    // Validate user input
    if err := services.ValidateStruct(user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "errors": helpers.FormatValidationError(err),
        })
    }

    if err := ctrl.UserService.CreateUser(user); err != nil {
        return c.Status(500).SendString("Failed to create user")
    }
    return c.JSON(user)
}

func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    user, err := ctrl.UserService.GetUserByID(uint(id))
    if err != nil {
        return c.Status(404).SendString("User not found")
    }

    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString("Invalid request payload")
    }

    // Validate updated user input
    if err := services.ValidateStruct(user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "errors": helpers.FormatValidationError(err),
        })
    }

    if err := ctrl.UserService.UpdateUser(user); err != nil {
        return c.Status(500).SendString("Failed to update user")
    }
    return c.JSON(user)
}

func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    if err := ctrl.UserService.DeleteUser(uint(id)); err != nil {
        return c.Status(500).SendString("Failed to delete user")
    }
    return c.SendStatus(204)
}
