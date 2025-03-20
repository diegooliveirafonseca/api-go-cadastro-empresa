package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/models"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secreto")

func Login(w http.ResponseWriter, r *http.Request) {
	// Receber dados de login (usuário e senha)
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
		return
	}

	// Buscar usuário no banco de dados
	user, err := models.GetUserByUsername(credentials.Username)
	if err != nil || user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	// Comparar senha fornecida com o hash armazenado no banco de dados
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Senha incorreta", http.StatusUnauthorized)
		return
	}

	// Criar o token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	// Responder com o token
	response := map[string]interface{}{
		"message": "Login bem-sucedido",
		"token":   tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
