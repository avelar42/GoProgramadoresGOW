package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname)

	var err error

	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Printf("Erro ao parsear config: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		// configurações de pool (ajuste conforme necessário)
		cfg.MaxConns = 500
		cfg.MinConns = 100
		cfg.MaxConnLifetime = 5 * time.Minute

		Pool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			log.Printf("Erro ao criar pool: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if err = Pool.Ping(ctx); err == nil {
			log.Println("✅ Banco conectado com pgxpool")
			return
		}

		log.Printf("Erro ao pingar DB: %v", err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("❌ Não foi possível conectar ao banco após várias tentativas: %v", err)
}
