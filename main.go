package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	})

	mux.HandleFunc("GET /home", getHome())

	mux.HandleFunc("GET /article/{id}", getArticle())

	mux.HandleFunc("GET /admin", basicAuth(getAdmin()))

	mux.HandleFunc("GET /edit/{id}", basicAuth(getEdit()))
	mux.HandleFunc("POST /edit/{id}", basicAuth(postEdit()))

	mux.HandleFunc("GET /new", basicAuth(getNew()))
	mux.HandleFunc("POST /new", basicAuth(postNew()))

	server := http.Server{
		Addr:    ":4000",
		Handler: logging(mux),
	}

	fmt.Printf("Server is up and running on port: %v\n", server.Addr[1:])
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
