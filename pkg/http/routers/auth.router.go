package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/middlewares"
	"github.com/gorilla/mux"
)

func RegisterAuthRouter(router *mux.Router, ctl *controllers.AuthCtl) {
	router.HandleFunc("/auth", ctl.Auth).Methods("POST")
	router.HandleFunc("/register", ctl.Register).Methods("POST")
	router.HandleFunc("/verifytoken", middlewares.ValidateToken(ctl.VerifyToken)).Methods("GET")
}
