package main

import (
	"crypto/subtle"
	"log"
	"net/http"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Println(err.Error())
		} else if cookie.Value == "session_token" {
			next.ServeHTTP(w, r)
			return
		}

		username, password, ok := r.BasicAuth()
		if ok {
			usernameMatch := (subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1)
			passwordMatch := (subtle.ConstantTimeCompare([]byte(password), []byte("admin")) == 1)

			if usernameMatch && passwordMatch {
				http.SetCookie(w, &http.Cookie{
					Name:     "session_token",
					Value:    "session_token",
					HttpOnly: true,
				})
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
