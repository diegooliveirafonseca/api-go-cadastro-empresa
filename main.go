package main

import (
	"fmt"
	"net/http"

	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/database"
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/routes"
)

func main() {
	database.Connect()
	r := routes.SetupRoutes()

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", r)
}
