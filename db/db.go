package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	for i := 0; i < 10; i++ {
		DB, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Erro ao abrir conexão: %v\n", err)
		} else {
			err = DB.Ping()
			if err == nil {
				fmt.Println("✅ Banco conectado")
				return
			}
			log.Printf("Erro ao pingar DB: %v\n", err)
		}
		log.Println("Tentativa de conexão falhou, tentando novamente em 2 segundos...")
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Não foi possível conectar ao banco após várias tentativas: %v", err)
}
