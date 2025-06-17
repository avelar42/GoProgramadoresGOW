package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()
    app.Get("/check", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "status":  "200",
            "message": "Servidor rodando com Go-Fiber",
        })
    })
    app.Listen(":8000")
}