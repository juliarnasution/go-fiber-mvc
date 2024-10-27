package controllers

import (
    "github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func NewHomeController() *HomeController {
    return &HomeController{}
}

// RenderHomePage renders a simple view
func (ctrl *HomeController) RenderHomePage(c *fiber.Ctx) error {
    data := fiber.Map{
        "Name":  "John Doe",
        "Email": "john.doe@example.com",
    }
    return c.Render("index", data)
}
