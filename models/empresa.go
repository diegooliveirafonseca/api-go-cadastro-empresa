package models

import (
	"database/sql"
	"fmt"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/database"
)

type Empresa struct {
	ID   int    `json:"id"`
	CNPJ string `json:"cnpj"`
	Nome string `json:"nome"`
}

// CriarEmpresa insere uma nova empresa no banco e retorna o ID gerado
func CriarEmpresa(empresa *Empresa) error {
	query := "INSERT INTO empresas (cnpj, nome) VALUES ($1, $2) RETURNING id"

	// Agora utilizamos database.DB
	err := database.DB.QueryRow(query, empresa.CNPJ, empresa.Nome).Scan(&empresa.ID)
	if err != nil {
		fmt.Println("Erro ao inserir empresa:", err)
		return err
	}
	fmt.Println("Empresa cadastrada com ID:", empresa.ID)
	return nil
}

func GetEmpresas() ([]Empresa, error) {
	query := "SELECT id, cnpj, nome FROM empresas"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar empresas: %v", err)
	}
	defer rows.Close()

	var empresas []Empresa
	for rows.Next() {
		var e Empresa
		if err := rows.Scan(&e.ID, &e.CNPJ, &e.Nome); err != nil {
			return nil, fmt.Errorf("erro ao escanear linha: %v", err)
		}
		empresas = append(empresas, e)
	}

	return empresas, nil
}

// GetEmpresaByCNPJ consulta uma empresa pelo CNPJ
func GetEmpresaByCNPJ(cnpj string) (*Empresa, error) {
	var empresa Empresa
	query := "SELECT id, cnpj, nome FROM empresas WHERE cnpj = $1"
	err := database.DB.QueryRow(query, cnpj).Scan(&empresa.ID, &empresa.CNPJ, &empresa.Nome)

	if err != nil {
		// Se o erro for sql.ErrNoRows, retornamos nil, nil para indicar que não encontrou a empresa
		if err == sql.ErrNoRows {
			// Aqui retorna nil, nil, o que é tratado como "empresa não encontrada"
			return nil, nil
		}
		// Qualquer outro erro retornamos como erro genérico
		return nil, err
	}

	return &empresa, nil
}

// DeleteEmpresa deleta uma empresa pelo CNPJ
func DeletarEmpresa(cnpj string) error {
	query := "DELETE FROM empresas WHERE cnpj = $1"
	result, err := database.DB.Exec(query, cnpj)
	if err != nil {
		return fmt.Errorf("erro ao deletar empresa: %v", err)
	}

	// Verifica se alguma linha foi afetada
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar exclusão: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma empresa encontrada com o CNPJ %s", cnpj)
	}

	return nil
}

func EmpresaExiste(cnpj string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM empresas WHERE cnpj = $1)"
	err := database.DB.QueryRow(query, cnpj).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar existência da empresa: %v", err)
	}
	return exists, nil
}
