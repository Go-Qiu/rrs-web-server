package routers

// New instantiate an instance a gorilla/mux router that contains the mapping of endpoints and handlers.
// func New(ctl *application.AppControllers) *mux.Router {

// 	// instantiate a gorilla/mux reouter.
// 	r := mux.NewRouter()
// 	PUBLIC := os.Getenv("PUBLIC")

// 	r.HandleFunc("/", handlers.ServeHtmlIndex)
// 	r.HandleFunc("/login", handlers.ServeHtmlLogin)
// 	r.HandleFunc("/logout", handlers.ServeHtmlLogout)

// 	// jwt auth api routes
// 	apiRouter := r.PathPrefix("/api/v1").Subrouter()
// 	RegisterAuthRouter(apiRouter, ctl.Auth)

// 	// users routes
// 	r.HandleFunc("/users", handlers.ServeHtmlIndexUsers)
// 	r.HandleFunc("/users/registration", handlers.ServeHtmlLogin)
// 	r.HandleFunc("/users/{id}", handlers.ServeHtmlUserProfile)
// 	r.HandleFunc("/users/{id}/transactions", handlers.ServeHtmlUserRecyclableTransactions)
// 	r.HandleFunc("/users/{id}/vouchers", handlers.ServeHtmlUserVouchers)
// 	r.HandleFunc("/users/{id}/points_to_vouchers/redepmtion", handlers.ServeHtmlUserPointsToVouchers)

// 	// vouchers routes
// 	r.HandleFunc("/vouchers", handlers.ServeHtmlIndexVouchers)

// 	// merchants routes
// 	r.HandleFunc("/merchants", handlers.ServeHtmlIndexMerchants)
// 	r.HandleFunc("/merchants/{id}/vouchers/aquire", handlers.ServeHtmlMerchantVouchersAquisition)
// 	r.HandleFunc("/merchants/{id}/vouchers/capture", handlers.ServeHtmlMerchantVoucherCapture)

// 	// static web pages or assets router
// 	fp := http.FileServer(http.Dir(PUBLIC))
// 	r.PathPrefix("/public").Handler(http.StripPrefix("/public/", fp))
// 	return r
// }

// RegisterRoutes bind children routes to a parent route.
// func RegisterHandlersWithRouter(router *mux.Router, handlers map[string]func(w http.ResponseWriter, r *http.Request)) {

// 	router.HandleFunc("/", handlers["/"]).Methods("POST")
// 	router.HandleFunc("/verifytoken", handlers["/verifytoken"]).Methods("GET")
// }
