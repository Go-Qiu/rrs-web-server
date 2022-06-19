package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// HandlesRoot will ....
func RequestForRoot(w http.ResponseWriter, r *http.Request) {

	// set the static web pages and assets folder path.
	PUBLIC := os.Getenv("PUBLIC")

	// root of web portal
	if r.URL.Path == "/" {
		ServeHtmlIndex(w, r)
		return
	}

	// request for static web pages or assets.
	// serve them.
	fp := filepath.Join(PUBLIC, filepath.Clean(r.URL.Path))
	http.ServeFile(w, r, fp)
}
