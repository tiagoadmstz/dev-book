package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Routes represents all web application routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.Response)
	RequiresAuthentication bool
}

//Configure puts all routes on the router
func Configure(router *mux.Router) *mux.Router {

}
