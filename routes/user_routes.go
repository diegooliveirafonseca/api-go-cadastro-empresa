package routes

import (
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/controllers"
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/middleware"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	secure := r.PathPrefix("/user").Subrouter()
	secure.Use(middleware.AuthMiddleware)

	secure.HandleFunc("/listarUsers", controllers.GetAllUsers).Methods("GET")
}
