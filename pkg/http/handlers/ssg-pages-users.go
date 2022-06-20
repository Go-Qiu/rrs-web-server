package handlers

import (
	"html/template"
	"net/http"
	"os"
)

func ServeHtmlUserRegistration(w http.ResponseWriter, r *http.Request) {

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
		Title:   "Web Portal - User Registration Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

// ServeHtmlUserProfile shows the web page containing (of a specific user), all the user details, summary vouchers info, summary points info, summary recyclable transactions info.  Also serve as the entry point to more details (vouchers, points, recyclable transactions) of the specific user.
func ServeHtmlUserProfile(w http.ResponseWriter, r *http.Request) {

	STATION_CODE := os.Getenv("STATION_CODE")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/header.tmpl.html",
		"./pkg/http/templates/users/profile.tmpl.html",
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
		Title:   "Web Portal - User Profile Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

// ServeHemlUserRecyclableTransactions will generate a web page that shows all the transactions info of a specific user.
func ServeHtmlUserRecyclableTransactions(w http.ResponseWriter, r *http.Request) {

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
		Title:   "Web Portal - User Recyclable Transactions Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

// ServeHtmlUserVouchers will generate a web page that shows all the active vouchers that a specific user has (by default).  It will have an option to list all the vouchers that he has consumed.
func ServeHtmlUserVouchers(w http.ResponseWriter, r *http.Request) {

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
		Title:   "Web Portal - User's Vouchers Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

// ServeHtmlUserPointsToVouchers will generate a web page that allows a user to convert his reward points (in hand) to vouchers (quantum and quantity he specify).
func ServeHtmlUserPointsToVouchers(w http.ResponseWriter, r *http.Request) {

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
		Title:   "Web Portal - Points to Vouchers Redepmtion Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}
