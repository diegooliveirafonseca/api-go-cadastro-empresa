package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5434"),
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

func CreateEmpresa(empresa *models.Empresa) error {
	// Query para inserir a empresa no banco
	query := `INSERT INTO empresas (nome, cnpj) VALUES ($1, $2, $3)`
	_, err := DB.Exec(query, empresa.Nome, empresa.CNPJ)
	if err != nil {
		return fmt.Errorf("erro ao criar empresa: %v", err)
	}
	return nil
}

func DeleteEmpresa(cnpj string) error {
	// Query para deletar a empresa
	query := `DELETE FROM empresas WHERE cnpj = $1`
	_, err := DB.Exec(query, cnpj)
	if err != nil {
		return fmt.Errorf("erro ao deletar empresa: %v", err)
	}
	return nil
}
