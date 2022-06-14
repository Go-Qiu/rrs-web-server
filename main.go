package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// get .env values
	err := godotenv.Load()
	if err != nil {
		errString := "[RRS]: fail to load .env"
		log.Fatal(errString)
	}
	SERVER_ADDR := os.Getenv("SERVER_ADDR")
	addr := SERVER_ADDR

	// start http server
	log.Printf("HTTPS Server started and listening on https://%s ...\n", addr)
	log.Fatalln(http.ListenAndServeTLS(addr, "./ssl/cert03.pem", "./ssl/key03.pem", nil))

}
