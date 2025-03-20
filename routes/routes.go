package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	AuthRoutes(r)
	UserRoutes(r)
	EmpresaRoutes(r)
	return r
}
