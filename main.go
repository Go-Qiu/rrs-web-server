package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-qiu/rrs-web-server/pkg/application"
	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/http/routers"
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

	// set the custom router.
	router := routers.New()

	// instantiate a controllers.
	authCtl := controllers.NewAuthCtl("Auth-SingPass", "Singpass-key")

	app := application.New()
	app.Controllers.Auth = authCtl
	app.Router = router

	// instantiate a custom http server.
	srv := http.Server{
		Addr:    SERVER_ADDR,
		Handler: app.Router,
	}

	// start http server.
	log.Printf("HTTPS Server started and listening on https://%s ...\n", SERVER_ADDR)
	log.Fatalln(srv.ListenAndServeTLS("./ssl/localhost.cert.pem", "./ssl/localhost.key.pem"))
	//
}
