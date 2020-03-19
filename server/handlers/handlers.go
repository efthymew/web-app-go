package handlers

import (
	"net/http"
)

//Index returns function to index
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/public/index.html")
}
