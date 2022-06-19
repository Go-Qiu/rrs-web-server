package handlers

import (
	"html/template"
	"net/http"
	"os"
)

// ServeHtmlIndex handles request for the home page.
func ServeHtmlIndex(w http.ResponseWriter, r *http.Request) {

	// get .env values
	// err := godotenv.Load()
	// if err != nil {
	// 	errString := "[JWT]: fail to load .env"
	// 	http.Error(w, errString, http.StatusInternalServerError)
	// 	return
	// }
	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/header.tmpl.html",
		"./pkg/http/templates/stations/index.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// app.serverError(w, err)
		return
	}

	// not a 'POST' request.
	// data to drive the template.
	tplData := struct {
		Station string
		Title   string
	}{
		Station: STATION_CODE,
		Title:   "HOME - Web Portal",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

func ServeHtmlIndexUsers(w http.ResponseWriter, r *http.Request) {

	// get .env values
	// err := godotenv.Load()
	// if err != nil {
	// 	errString := "[JWT]: fail to load .env"
	// 	http.Error(w, errString, http.StatusInternalServerError)
	// 	return
	// }
	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/stations/header.tmpl.html",
		"./pkg/http/templates/stations/index.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// app.serverError(w, err)
		return
	}

	// not a 'POST' request.
	// data to drive the template.
	tplData := struct {
		Station string
		Title   string
	}{
		Station: STATION_CODE,
		Title:   "HOME - Users",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

func ServeHtmlIndexVouchers(w http.ResponseWriter, r *http.Request) {

	// get .env values
	// err := godotenv.Load()
	// if err != nil {
	// 	errString := "[JWT]: fail to load .env"
	// 	http.Error(w, errString, http.StatusInternalServerError)
	// 	return
	// }
	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/stations/header.tmpl.html",
		"./pkg/http/templates/stations/index.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// app.serverError(w, err)
		return
	}

	// not a 'POST' request.
	// data to drive the template.
	tplData := struct {
		Station string
		Title   string
	}{
		Station: STATION_CODE,
		Title:   "HOME - Vouchers",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

func ServeHtmlIndexMerchants(w http.ResponseWriter, r *http.Request) {

	// get .env values
	// err := godotenv.Load()
	// if err != nil {
	// 	errString := "[JWT]: fail to load .env"
	// 	http.Error(w, errString, http.StatusInternalServerError)
	// 	return
	// }
	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/stations/header.tmpl.html",
		"./pkg/http/templates/stations/index.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// app.serverError(w, err)
		return
	}

	// not a 'POST' request.
	// data to drive the template.
	tplData := struct {
		Station string
		Title   string
	}{
		Station: STATION_CODE,
		Title:   "HOME - Merchants",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}
