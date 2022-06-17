package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/http/handlers"
	"github.com/gorilla/mux"
)

// New instantiate an instance a gorilla/mux router that contains the mapping of endpoints and handlers.
func New() *mux.Router {

	// instantiate a gorilla/mux reouter.
	r := mux.NewRouter()

	// api routesrs
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	usersAPIRouter := apiRouter.PathPrefix("/users").Subrouter()
	vouchersAPIRouter := apiRouter.PathPrefix("/vouchers").Subrouter()
	merchantsAPIRouter := apiRouter.PathPrefix("/merchants").Subrouter()

	RegisterAuthRouter(apiRouter, controllers.NewAuthCtl("auth", "singpass-api-key"))
	RegisterUsersRouter(usersAPIRouter, controllers.NewUsersCtl("users", "users-api-key"))
	RegisterVouchersRouter(vouchersAPIRouter, controllers.NewVouchersCtl("vouchers", "vouchers-api-key"))
	RegisterMerchantsRouter(merchantsAPIRouter, controllers.NewMerchantsCtl("merchants", "merchant-api-key"))

	// server-side-generated pages routers.

	// login
	// r.HandleFunc("/", handlers.ServeHtmlIndex)
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

	return r
}

// RegisterRoutes bind children routes to a parent route.
// func RegisterHandlersWithRouter(router *mux.Router, handlers map[string]func(w http.ResponseWriter, r *http.Request)) {

// 	router.HandleFunc("/", handlers["/"]).Methods("POST")
// 	router.HandleFunc("/verifytoken", handlers["/verifytoken"]).Methods("GET")
// }

// func Auth(w http.ResponseWriter, r *http.Request) {

// }

// func VerifyToken(w http.ResponseWriter, r *http.Request) {

// }
