package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"personal_blog/templates"
	"personal_blog/types"
	"strconv"
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
		var article types.Article
		var articles []types.Article

		entries, err := os.ReadDir("storage")
		if err != nil {
			log.Println(err.Error())
		}
		for i := len(entries) - 1; i >= 0; i-- {
			file, err := os.Open("./storage/" + entries[i].Name())
			if err != nil {
				log.Println(err.Error())
			}
			data, err := io.ReadAll(file)
			if err != nil {
				log.Println(err.Error())
			}
			json.Unmarshal(data, &article)
			articles = append(articles, article)
		}

		componant := templates.Admin(articles)
		componant.Render(context.Background(), w)
	}
}

func getEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article types.Article
		var articles []types.Article

		entries, err := os.ReadDir("storage")
		if err != nil {
			log.Println(err.Error())
		}
		for _, v := range entries {
			file, err := os.Open("./storage/" + v.Name())
			if err != nil {
				log.Println(err.Error())
			}
			data, err := io.ReadAll(file)
			if err != nil {
				log.Println(err.Error())
			}
			json.Unmarshal(data, &article)
			articles = append(articles, article)
		}

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

func postNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article types.Article

		_, err := os.Stat("./storage")
		if os.IsNotExist(err) {
			err := os.Mkdir("storage", 1644)
			if err != nil {
				log.Println(err.Error())
			}
		} else {
			entries, err := os.ReadDir("storage")
			if err != nil {
				log.Println(err.Error())
			}
			lastIndex := len(entries) - 1
			if len(entries) > 0 {
				article.ID, _ = strconv.Atoi(entries[lastIndex].Name()[:len(entries[lastIndex].Name())-5])
				article.ID++
			} else {
				article.ID = 1
			}
		}

		r.ParseForm()
		article.Title = r.FormValue("title")
		article.Date = r.FormValue("publish_date")
		article.Content = r.FormValue("content")

		jsonData, err := json.MarshalIndent(&article, "", "\t")
		if err != nil {
			log.Println(err.Error())
		}

		fileName := fmt.Sprintf("./storage/%v.json", article.ID)
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
		}

		file.Write(jsonData)

		http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
	}
}

func postEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "admin", http.StatusMovedPermanently)
	}
}
