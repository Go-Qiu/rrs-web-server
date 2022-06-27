package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterUsersRouter(router *mux.Router, ctl *controllers.UserCtl) {
	// router.HandleFunc("/{id}/transactions/type/{type_code}", ctl.GetTransactionsByType).Methods("GET")
	router.HandleFunc("/{id}/transactions", ctl.GetTransactionsByType).Methods("GET").Queries("type", "{type_code}")
	router.HandleFunc("/{id}/points", ctl.GetPoints)
}
