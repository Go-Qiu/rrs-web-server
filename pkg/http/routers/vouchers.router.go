package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterVouchersRouter(router *mux.Router, ctl *controllers.VouchersCtl) {
	router.HandleFunc("/", ctl.GetAll).Methods("GET")
}
