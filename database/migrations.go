package database

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func RunMigrations() {
	migrationQueries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(100) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS empresas (
			id SERIAL PRIMARY KEY,
			cnpj VARCHAR(14) UNIQUE NOT NULL,
			nome VARCHAR(255) NOT NULL
		);`,
	}

	for _, query := range migrationQueries {
		_, err := DB.Exec(query)
		if err != nil {
			fmt.Println("Erro ao executar migration:", err)
		} else {
			fmt.Println("Migration aplicada com sucesso!")
		}
	}

	SeedData()
}

func SeedData() {
	seedEmpresas()
	seedUsers()
}

func seedEmpresas() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM empresas").Scan(&count)
	if err != nil {
		fmt.Println("Erro ao verificar dados iniciais das empresas:", err)
		return
	}

	if count == 0 {
		fmt.Println("Inserindo dados iniciais nas empresas...")

		_, err = DB.Exec(`INSERT INTO empresas (cnpj, nome) VALUES
			('12345678000101', 'Tech Solutions'),
			('98765432000199', 'Market Digital'),
			('12345678901235', 'Empresa 2'),
			('12345678901237', 'Nova Empresa 2');`)

		if err != nil {
			fmt.Println("Erro ao inserir dados iniciais nas empresas:", err)
		} else {
			fmt.Println("Dados iniciais de empresas inseridos com sucesso!")
		}
	}
}

func seedUsers() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		fmt.Println("Erro ao verificar dados iniciais dos usuários:", err)
		return
	}

	if count == 0 {
		fmt.Println("Inserindo dados iniciais nos usuários...")

		// Hash de senha para segurança
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Erro ao gerar hash da senha:", err)
		}

		_, err = DB.Exec(`INSERT INTO users (username, password_hash, role) VALUES
			('admin', $1, 'admin');`, string(hashedPassword))

		if err != nil {
			fmt.Println("Erro ao inserir dados iniciais nos usuários:", err)
		} else {
			fmt.Println("Dados iniciais de usuários inseridos com sucesso!")
		}
	}
}
