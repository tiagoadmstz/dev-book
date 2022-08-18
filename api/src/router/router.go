package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate return route configuration
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
