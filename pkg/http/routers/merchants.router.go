package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterMerchantsRouter(router *mux.Router, ctl *controllers.MerchantsCtl) {
	router.HandleFunc("/", ctl.GetAll).Methods("GET")
}
