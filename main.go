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
