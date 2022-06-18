package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterMerchantsRouter(router *mux.Router, ctl *controllers.MerchantsCtl) {
	router.HandleFunc("/", ctl.GetAll).Methods("GET")
	router.HandleFunc("/", ctl.Create).Methods("POST")
	router.HandleFunc("/{id}", ctl.GetById).Methods("GET")
	router.HandleFunc("/{id}", ctl.HandlePutRequest).Methods("PUT")
	router.HandleFunc("/{id}/branches", ctl.AddBranches).Methods("POST")
	router.HandleFunc("/{id}/branches", ctl.RemoveBranches).Methods("DELETE")
}
