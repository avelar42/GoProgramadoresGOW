package controllers

import (
	"github.com/avelar42/GoProgramadoresGOW/db"
	"github.com/avelar42/GoProgramadoresGOW/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	if len(p.Apelido) > 32 || len(p.Nome) > 100 {
		return c.Status(422).SendString("Campos muito longos")
	}

	id := uuid.New().String()

	query := `
	INSERT INTO programadores (id, apelido, nome, nascimento, stack)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.DB.Exec(query, id, p.Apelido, p.Nome, p.Nascimento, pq.Array(p.Stack))
	if err != nil {
		return c.Status(422).SendString("Erro ao inserir: " + err.Error())
	}

	c.Location("/programadores/" + id)
	return c.SendStatus(201)
}
