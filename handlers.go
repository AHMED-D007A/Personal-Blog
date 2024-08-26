package main

import (
	"context"
	"net/http"
	"personal_blog/templates"
)

func getHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Home()
		componant.Render(context.Background(), w)
	}
}

func getArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Article()
		componant.Render(context.Background(), w)
	}
}

func getAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Admin()
		componant.Render(context.Background(), w)
	}
}

func getEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.Edit()
		componant.Render(context.Background(), w)
	}
}

func getNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		componant := templates.New()
		componant.Render(context.Background(), w)
	}
}
