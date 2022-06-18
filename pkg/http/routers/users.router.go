package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterUsersRouter(router *mux.Router, ctl *controllers.UsersCtl) {
	router.HandleFunc("/", ctl.GetAll).Methods("GET")
	router.HandleFunc("/{id}", ctl.GetById).Methods("GET")
	router.HandleFunc("/{id}", ctl.Inactivate).Methods("PUT")
	router.HandleFunc("/{id}/transactions", ctl.AddTransaction).Methods("POST")
	router.HandleFunc("/{id}/xchange", ctl.XChangePointsToVouchers).Methods("POST")
	router.HandleFunc("/{id}/redeem/vouchers", ctl.RedeemVouchers).Methods("POST")

}
