package handlers

import (
	"html/template"
	"net/http"
	"os"
)

func ServeHtmlMerchantVouchersAquired(w http.ResponseWriter, r *http.Request) {

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
		Title:   "Web Portal - Merchant Acquire Vouchers Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}

func ServeHtmlMerchantVoucherCapture(w http.ResponseWriter, r *http.Request) {

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
		Title:   "Web Portal - Merchant Voucher Capture Page",
	}

	ts.ExecuteTemplate(w, "base", tplData)
}
