package routes

import "github.com/gorilla/mux"

// NewStationRoutes instantiate an instance of the station routes.
func NewStationRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex).Methods("Get")
	r.HandleFunc("/dropoff", handleDropOff).Methods("GET")
	r.HandleFunc("/dropoff", handleItemsToPoints).Methods("POST")

	return r
}
