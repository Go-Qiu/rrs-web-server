package routes

import "github.com/gorilla/mux"

// New instantiate an instance a gorilla/mux router that contains the mapping of endpoints and handlers.
func New() *mux.Router {

	// instantiate a gorilla/mux reouter.
	r := mux.NewRouter()

	// register all the defined routes.
	generic := r.PathPrefix("/api/v1").Subrouter()

	// authentication routes

	// users routes
	usersRoutes := generic.PathPrefix("/users").Subrouter()

	// vouchers routes

	// merchants routes

	// station routes
	return r
}

// RegisterRoutes bind children routes to a parent route.
func RegisterRoutes(parent *mux.Router, child *mux.Router) {

}
