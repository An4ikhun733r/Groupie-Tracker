package web

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", Artist)
	mux.HandleFunc("/search", Search)
	mux.HandleFunc("/filter", Filter)
	mux.HandleFunc("/autocomplete", Autocomplete)

	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on http://localhost:4000")

	http.ListenAndServe(":8000", mux)
}
