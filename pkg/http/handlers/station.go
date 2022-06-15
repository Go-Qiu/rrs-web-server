package handlers

import (
	"html/template"
	"net/http"
)

// HandleIndex handles the server-side-generation of the Station index page.
func ServerHtmlIndex(w http.ResponseWriter, r *http.Request) {

	// set the response header attribute, "Content-Type" to "text/html".  Indicates the response is a HTML page.
	w.Header().Set("Content-Type", "text/html")

	files := []string{
		"./pkg/http/templates/base.tmpl.html",
		"./pkg/http/templates/header.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// more code in passing data into template

	// render the template
	ts.ExecuteTemplate(w, "base", nil)

}
