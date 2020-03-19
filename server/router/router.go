package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Router returns new mux router with endpoints
func Router() *mux.Router {
	r := mux.NewRouter()
	//serve static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/build"))))
	//r.PathPrefix("/").HandlerFunc(handlers.Index)
	//http.Handle("/", r)
	return r
}
