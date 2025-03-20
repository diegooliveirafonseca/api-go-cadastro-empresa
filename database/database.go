package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres123"),
		getEnv("DB_NAME", "goapi_development"),
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	fmt.Println("Banco de dados conectado!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
