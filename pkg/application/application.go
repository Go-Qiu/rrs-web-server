package application

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/gorilla/mux"
)

type AppControllers struct {
	CRUD AppCRUDControllers
	Auth *controllers.AuthCtl
}

type AppCRUDControllers struct {
	Users     *controllers.CRUD
	Vouchers  *controllers.CRUD
	Merchants *controllers.CRUD
}

type Application struct {
	Controllers AppControllers
	Router      *mux.Router
}

func New() *Application {
	return &Application{}
}
