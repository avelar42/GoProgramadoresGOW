package main

import (
	"log"
	"github.com/avelar42/GoProgramadoresGOW/db"
	"github.com/avelar42/GoProgramadoresGOW/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db.Connect()            // conecta ao PostgreSQL
	routes.SetupRoutes(app) // configura rotas

	log.Fatal(app.Listen(":3000"))
}