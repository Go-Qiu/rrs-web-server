package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterStationRouter(router *mux.Router, ctl *controllers.TransactionCtl) {
	router.HandleFunc("/{id}/transactions", ctl.Create).Methods("POST")
}
