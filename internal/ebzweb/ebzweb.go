package ebzweb

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:public
var web embed.FS

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	index, err := web.ReadFile("public/index.html")
	if err != nil {
		http.Error(rw, "Not Found", http.StatusNotFound)
		return
	}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.Write(index)
}

func New(mux *http.ServeMux) *http.ServeMux {
	publicFS, _ := fs.Sub(web, "public")
	mux.Handle("GET /assets/", http.FileServer(http.FS(publicFS)))
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /vite.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, publicFS, "vite.svg")
	})
	return mux
}
