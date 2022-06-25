package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-qiu/rrs-web-server/pkg/application"
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/joho/godotenv"
)

func main() {

	// get .env values
	err := godotenv.Load()
	if err != nil {

		// fail
		errString := "[RRS]: fail to load .env"
		log.Fatal(errString)
	}
	SERVER_ADDR := os.Getenv("SERVER_ADDR")
	// API_KEY_USERS := os.Getenv("API_KEY_USERS")
	// API_KEY_VOUCHERS := os.Getenv("API_KEY_VOUCHERS")
	// API_KEY_MERCHANTS := os.Getenv("API_KEY_MERCHANTS")

	// set the custom router.
	// router := routers.New()

	jwtConfig := controllers.JWTConfig{
		ISSUER:     os.Getenv("JWT_ISSUER"),
		EXP_MIN:    os.Getenv("JWT_EXP_MINUTES"),
		SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
	}

	// instantiate an in-memory data store, to cache the Users data points.
	// var dsUsers ds.DataStore = *ds.New()

	// instantiate an application to link the respective controllers and router.
	app := application.New()
	app.DataStore = make(map[string]controllers.DataPointExtended)

	// populate the users in-memory cache, using the users data from users microservice.
	app.PullDataIntoDataStore()

	// instantiate a authentication controller.
	authCtl := controllers.NewAuthCtl("JWT AUTH SERVICE", "", &jwtConfig, app.DataStore)
	app.Controllers.Auth = authCtl

	// instantiate a voucher controller.

	// instantiate a merchant controller.

	// app.Router = router

	// instantiate a custom http server.
	srv := http.Server{
		Addr:    SERVER_ADDR,
		Handler: app.Router(),
	}

	// start http server.
	log.Printf("HTTPS Server started and listening on https://%s ...\n", SERVER_ADDR)
	log.Fatalln(srv.ListenAndServeTLS("./ssl/localhost.cert.pem", "./ssl/localhost.key.pem"))
	//
}
