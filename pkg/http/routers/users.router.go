package routers

import (
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/middlewares"
	"github.com/gorilla/mux"
)

func RegisterUsersRouter(router *mux.Router, ctl *controllers.UserCtl) {

	// route to get all transactions by item type, for a specific user.
	// type_code --> "PLASTIC", "PAPER", "METAL", "GLASS".
	router.HandleFunc("/{id}/transactions", middlewares.ValidateToken(ctl.GetTransactionsByType)).Methods("GET").Queries("type", "{type_code}")

	// route to get the points collected by the user.
	router.HandleFunc("/{id}/points", middlewares.ValidateToken(ctl.GetPoints)).Methods("GET")

	// route to get the usable vouchers that a user has in possession.
	router.HandleFunc("/{id}/voucers", ctl.GetVouchers).Methods("GET")

	router.HandleFunc("/{id}/points_to/vouchers", ctl.PointsToVouchers).Methods("POST")
}
