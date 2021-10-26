package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8999",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err != nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
