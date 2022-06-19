package routers

import (
	"net/http"
	"os"

	"github.com/go-qiu/rrs-web-server/pkg/http/handlers"
	"github.com/gorilla/mux"
)

// New instantiate an instance a gorilla/mux router that contains the mapping of endpoints and handlers.
func New() *mux.Router {

	// instantiate a gorilla/mux reouter.
	r := mux.NewRouter()
	PUBLIC := os.Getenv("PUBLIC")

	// register all the defined routes.
	// parentGeneric := r.PathPrefix("/api/v1").Subrouter()

	// authentication routes
	// handlers := make(map[string]func(w http.ResponseWriter, r *http.Request))
	// handlers["/"] = Auth
	// handlers["/verifytoken"] = VerifyToken

	// RegisterHandlersWithRouter(parentGeneric, handlers)

	// login

	r.HandleFunc("/", handlers.RequestForRoot)
	// r.HandleFunc("/login", handlers.ServeHtmlLogin)

	// users routes
	r.HandleFunc("/users", handlers.ServeHtmlIndexUsers)

	// vouchers routes
	r.HandleFunc("/vouchers", handlers.ServeHtmlIndexVouchers)

	// r.HandleFunc("/vouchers/sponsors/", h)

	// merchants routes
	r.HandleFunc("/merchants", handlers.ServeHtmlIndexMerchants)

	// station routes
	// stationRoutes := generic.PathPrefix("/station").Subrouter()
	// stationRoutes.HandleFunc("/", handlers.ServerHtmlStationIndex)
	// stationRoutes.HandleFunc("/dropoff", handlers.ServerHtmlStationDropOff)

	// static web pages or assets router
	fp := http.FileServer(http.Dir(PUBLIC))
	r.PathPrefix("/public").Handler(http.StripPrefix("/public/", fp))
	return r
}

// RegisterRoutes bind children routes to a parent route.
func RegisterHandlersWithRouter(router *mux.Router, handlers map[string]func(w http.ResponseWriter, r *http.Request)) {

	router.HandleFunc("/", handlers["/"]).Methods("POST")
	router.HandleFunc("/verifytoken", handlers["/verifytoken"]).Methods("GET")
}

func Auth(w http.ResponseWriter, r *http.Request) {

}

func VerifyToken(w http.ResponseWriter, r *http.Request) {

}
