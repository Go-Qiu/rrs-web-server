package handlers

import (
	"html/template"
	"net/http"
	"os"
)

func ServeHtmlLogin(w http.ResponseWriter, r *http.Request) {

	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/header.tmpl.html",
		"./pkg/http/templates/login/login.tmpl.html",
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
		Title:   "Web Portal - Login",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

func ServeHtmlLogout(w http.ResponseWriter, r *http.Request) {

	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/header.tmpl.html",
		"./pkg/http/templates/login/logout.tmpl.html",
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
		Title:   "Web Portal - Logout Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}
