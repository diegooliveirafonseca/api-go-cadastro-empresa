package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/models"
	"golang.org/x/crypto/bcrypt"
)

// Definir uma struct para a resposta sem os campos sensíveis
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// RegisterUser cria um novo usuário no sistema
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar os dados", http.StatusBadRequest)
		return
	}

	// Gerar hash para a senha
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao gerar hash da senha", http.StatusInternalServerError)
		return
	}

	// Salvar o novo usuário no banco de dados
	err = models.CreateUser(user.Username, string(passwordHash), user.Role)
	if err != nil {
		http.Error(w, "Erro ao criar o usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso!"})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	// Criar uma lista filtrada de usuários
	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersResponse)
}
