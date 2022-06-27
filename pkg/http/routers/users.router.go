package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/middlewares"
	"github.com/gorilla/mux"
)

func RegisterUsersRouter(router *mux.Router, ctl *controllers.UserCtl) {
	// router.HandleFunc("/{id}/transactions/type/{type_code}", ctl.GetTransactionsByType).Methods("GET")
	router.HandleFunc("/{id}/transactions", middlewares.ValidateToken(ctl.GetTransactionsByType)).Methods("GET").Queries("type", "{type_code}")
	router.HandleFunc("/{id}/points", middlewares.ValidateToken(ctl.GetPoints)).Methods("GET")
	router.HandleFunc("/{id}/voucers", ctl.GetVouchers).Methods("GET")
}
