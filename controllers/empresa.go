package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/models"
)

// GetEmpresa consulta a empresa pelo CNPJ
func GetEmpresa(w http.ResponseWriter, r *http.Request) {
	// Captura o parâmetro 'cnpj' da query string
	cnpj := r.URL.Query().Get("cnpj")

	// Valida se o CNPJ foi fornecido
	if cnpj == "" {
		http.Error(w, "CNPJ é necessário", http.StatusBadRequest)
		return
	}

	// Chama a função para buscar a empresa
	empresa, err := models.GetEmpresaByCNPJ(cnpj)
	if err != nil {
		// Se houver erro, retornamos erro interno
		http.Error(w, "Erro ao consultar empresa", http.StatusInternalServerError)
		return
	}

	// Se empresa for nil, ou seja, não encontrada
	if empresa == nil {
		http.Error(w, fmt.Sprintf("Empresa com CNPJ %s não encontrada", cnpj), http.StatusNotFound)
		return
	}

	// Caso a empresa seja encontrada, retornamos os dados em JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresa)
}

func GetAllEmpresas(w http.ResponseWriter, r *http.Request) {
	empresas, err := models.GetEmpresas()
	if err != nil {
		http.Error(w, "Erro ao buscar empresas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresas)
}

// CreateEmpresa cria uma nova empresa
func CreateEmpresa(w http.ResponseWriter, r *http.Request) {
	var empresa models.Empresa

	// Decodificar o corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&empresa)
	if err != nil {
		http.Error(w, "Erro ao decodificar os dados", http.StatusBadRequest)
		return
	}

	// Criar a empresa no banco
	err = models.CriarEmpresa(&empresa)
	if err != nil {
		http.Error(w, "Erro ao criar empresa: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com sucesso
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Empresa criada com sucesso!"}
	json.NewEncoder(w).Encode(response)
}

// DeleteEmpresa deleta uma empresa pelo CNPJ
func DeleteEmpresa(w http.ResponseWriter, r *http.Request) {
	// Captura o parâmetro 'cnpj' da query string
	cnpj := r.URL.Query().Get("cnpj")

	// Valida se o CNPJ foi fornecido
	if cnpj == "" {
		http.Error(w, "CNPJ é necessário", http.StatusBadRequest)
		return
	}

	// Primeiro, verifica se a empresa existe
	existe, err := models.EmpresaExiste(cnpj)
	if err != nil {
		http.Error(w, "Erro ao verificar empresa", http.StatusInternalServerError)
		return
	}

	if !existe {
		http.Error(w, fmt.Sprintf("Empresa com CNPJ %s não encontrada", cnpj), http.StatusNotFound)
		return
	}

	// Chama a função para deletar a empresa
	err = models.DeletarEmpresa(cnpj)
	if err != nil {
		http.Error(w, "Erro ao deletar empresa: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com sucesso
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Empresa deletada com sucesso!"}
	json.NewEncoder(w).Encode(response)
}
