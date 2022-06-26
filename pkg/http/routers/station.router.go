package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/middlewares"
	"github.com/gorilla/mux"
)

func RegisterStationRouter(router *mux.Router, ctl *controllers.TransactionCtl) {
	router.HandleFunc("/{id}/transactions", middlewares.ValidateToken(ctl.Create)).Methods("POST")
}
