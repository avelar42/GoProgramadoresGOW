package routes

import (
	"github.com/avelar42/GoProgramadoresGOW/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/programadores", controllers.CriarProgramador)
}
