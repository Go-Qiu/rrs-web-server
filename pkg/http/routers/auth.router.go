package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterAuthRouter(router *mux.Router, ctl *controllers.AuthCtl) {
	router.HandleFunc("/auth", ctl.Auth).Methods("POST")
}
