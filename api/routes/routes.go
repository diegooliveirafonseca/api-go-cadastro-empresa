package routes

import (
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/controllers"
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/api/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")

	secure := r.PathPrefix("/empresa").Subrouter()
	secure.Use(middleware.AuthMiddleware)
	secure.HandleFunc("", controllers.GetEmpresa).Methods("GET")
	secure.HandleFunc("", controllers.CreateEmpresa).Methods("POST")
	secure.HandleFunc("", controllers.DeleteEmpresa).Methods("DELETE")

	return r
}
