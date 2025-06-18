package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5433 user=user password=Pass123$ dbname=Programadores sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao pingar DB:", err)
	}

	fmt.Println("✅ Banco conectado")
}
