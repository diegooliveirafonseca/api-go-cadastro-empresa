package models

import (
	"fmt"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/database"
)

// User representa o modelo de dados de um usuário
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"` // Não deve ser armazenado, apenas usado ao criar
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

// GetUsers retorna todos os usuários cadastrados no banco de dados
func GetUsers() ([]User, error) {
	query := "SELECT id, username, role FROM users"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
			return nil, fmt.Errorf("erro ao escanear linha: %v", err)
		}
		users = append(users, u)
	}

	return users, nil
}

// GetUserByUsername busca um usuário pelo nome de usuário
func GetUserByUsername(username string) (*User, error) {
	query := "SELECT id, username, password_hash, role FROM users WHERE username = $1"
	row := database.DB.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}

	return &user, nil
}

// CreateUser cria um novo usuário no banco de dados
func CreateUser(username, passwordHash, role string) error {
	query := "INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, username, passwordHash, role)
	if err != nil {
		return fmt.Errorf("erro ao criar usuário: %v", err)
	}

	return nil
}
