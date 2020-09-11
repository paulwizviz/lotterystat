package ebzmux

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed static/htmx.min.js templates/* css/main.css
var web embed.FS

func staticHandler(rw http.ResponseWriter, r *http.Request) {
	staticContent, err := web.ReadFile("static/htmx.min.js")
	if err != nil {
		http.NotFound(rw, r)
		return
	}
	rw.Header().Set("Content-Type", "application/javascript")
	_, err = rw.Write(staticContent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func cssHandler(rw http.ResponseWriter, r *http.Request) {
	cssContent, err := web.ReadFile("css/main.css")
	if err != nil {
		http.NotFound(rw, r)
		return
	}
	rw.Header().Set("Content-Type", "text/css")
	_, err = rw.Write(cssContent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("base").ParseFS(web, "templates/*.html"))
	data := map[string]string{"Title": "Home", "Welcome": "Hello, world!"}
	err := t.ExecuteTemplate(rw, "homePage", data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func clickedHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	_, err := rw.Write([]byte(`<strong>Button clicked!</strong>`))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /static/htmx.min.js", staticHandler)
	mux.HandleFunc("GET /css/main.css", cssHandler)
	mux.HandleFunc("GET /clicked", clickedHandler)
	return mux
}
