package routes

import (
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/controllers"
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/middleware"
	"github.com/gorilla/mux"
)

func EmpresaRoutes(r *mux.Router) {
	// Criando um subrouter seguro (exige autenticação)
	secure := r.PathPrefix("/empresa").Subrouter()
	secure.Use(middleware.AuthMiddleware)

	// Todas as operações só podem ser feitas por usuários autenticados
	secure.HandleFunc("/listarEmpresas", controllers.GetAllEmpresas).Methods("GET")
	secure.HandleFunc("", controllers.GetEmpresa).Methods("GET")
	secure.HandleFunc("/empresas", controllers.CreateEmpresa).Methods("POST")
	secure.HandleFunc("", controllers.DeleteEmpresa).Methods("DELETE")
}
