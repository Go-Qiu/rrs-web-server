package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ServeHtml(w http.ResponseWriter, r *http.Request) {
	// get .env values.
	// err := godotenv.Load()
	// if err != nil {
	// 	errString := "[WEB]: fail to load .env"
	// 	http.Error(w, errString, http.StatusInternalServerError)
	// 	return
	// }
	PUBLIC := os.Getenv("PUBLIC")
	MODE := os.Getenv("MODE")

	// handling when the mode is 'DEV'
	if MODE == "DEV" {
		log.Printf("[STATIC] received request for %s", r.URL.Path)
	}

	// request for home page.
	if r.URL.Path == "/" {
		ServeHtmlIndex(w, r)
		return
	}

	// serve the html page
	fp := filepath.Join(PUBLIC, filepath.Clean(r.URL.Path))
	http.ServeFile(w, r, fp)
}
