package routes

import (
	"github.com/diegooliveirafonseca/api-go-cadastro-empresa/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
