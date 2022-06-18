package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterVouchersRouter(router *mux.Router, ctl *controllers.VouchersCtl) {
	router.HandleFunc("/", ctl.GetAll).Methods("GET")
	router.HandleFunc("/", ctl.HandlePostRequest).Methods("POST")
	router.HandleFunc("/{id}", ctl.GetById).Methods("GET")
	router.HandleFunc("/{id}", ctl.RedeemById).Methods("PUT")

}
