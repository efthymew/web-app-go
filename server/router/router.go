package router

import (
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

//Router returns new mux router with endpoints
func Router() *mux.Router {
	r := mux.NewRouter()
	//serve static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", CustomFileServer(http.Dir("./frontend/build"))))
	//r.PathPrefix("/").HandlerFunc(handlers.Index)
	//http.Handle("/", r)
	return r
}

//CustomFileServer c
func CustomFileServer(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			NotFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

//NotFound n
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/build/error404.html")
}
