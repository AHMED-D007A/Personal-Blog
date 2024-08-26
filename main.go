package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal_blog/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	})

	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Home()
		componant.Render(context.Background(), w)
	})

	mux.HandleFunc("GET /article/{id}", func(w http.ResponseWriter, r *http.Request) {
	})

	mux.HandleFunc("GET /admin", func(w http.ResponseWriter, r *http.Request) {
	})

	mux.HandleFunc("GET /edit/{id}", func(w http.ResponseWriter, r *http.Request) {
	})

	mux.HandleFunc("GET /new", func(w http.ResponseWriter, r *http.Request) {
	})

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
