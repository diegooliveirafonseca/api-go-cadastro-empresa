package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/database"
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/models"
	"github.com/gorilla/mux"
)

var empresas = []models.Empresa{}

func GetEmpresa(w http.ResponseWriter, r *http.Request) {
	cnpj := r.URL.Query().Get("cnpj")
	for _, empresa := range empresas {
		if empresa.CNPJ == cnpj {
			json.NewEncoder(w).Encode(empresa)
			return
		}
	}
	http.Error(w, "Empresa não encontrada", http.StatusNotFound)
}

func CreateEmpresa(w http.ResponseWriter, r *http.Request) {
	var empresa models.Empresa

	// Decodificar o corpo da requisição para a struct empresa
	err := json.NewDecoder(r.Body).Decode(&empresa)
	if err != nil {
		http.Error(w, "Erro ao decodificar os dados", http.StatusBadRequest)
		return
	}

	// Inserir a empresa no banco de dados
	err = database.CreateEmpresa(&empresa)
	if err != nil {
		http.Error(w, "Erro ao criar empresa: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com uma mensagem de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Empresa criada com sucesso!"}
	json.NewEncoder(w).Encode(response)
}

func DeleteEmpresa(w http.ResponseWriter, r *http.Request) {
	// Extrair o CNPJ da URL
	params := mux.Vars(r)
	cnpj := params["cnpj"]

	// Deletar a empresa do banco de dados
	err := database.DeleteEmpresa(cnpj)
	if err != nil {
		http.Error(w, "Erro ao deletar empresa: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com uma mensagem de sucesso
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Empresa deletada com sucesso!"}
	json.NewEncoder(w).Encode(response)
}
