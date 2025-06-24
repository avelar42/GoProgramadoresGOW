package controllers

import (
	"context"

	"github.com/avelar42/GoProgramadoresGOW/db"
	"github.com/avelar42/GoProgramadoresGOW/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CriarProgramador(c *fiber.Ctx) error {
	var p models.Programador

	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).SendString("JSON inválido")
	}

	// Validações
	if p.Apelido == "" || p.Nome == "" || p.Nascimento == "" {
		return c.Status(422).SendString("Campos obrigatórios ausentes")
	}
	if len(p.Apelido) > 100 || len(p.Nome) > 100 {
		return c.Status(422).SendString("Campos muito longos")
	}

	id := uuid.New().String()

	query := `
		INSERT INTO programadores (id, apelido, nome, nascimento, stack)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Pool.Exec(
		context.Background(),
		query,
		id, p.Apelido, p.Nome, p.Nascimento, p.Stack,
	)
	if err != nil {
		return c.Status(422).SendString("Erro ao inserir: " + err.Error())
	}

	c.Location("/programadores/" + id)
	return c.SendStatus(201)
}

func ContarProgramadores(c *fiber.Ctx) error {
	var count int

	err := db.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM programadores").Scan(&count)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao contar programadores",
		})
	}

	return c.JSON(fiber.Map{
		"total": count,
	})
}
