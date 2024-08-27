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
		var article types.Article
		var articles []types.Article
		var file *os.File

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

		defer file.Close()

		componant := templates.Home(articles)
		componant.Render(context.Background(), w)
	}
}

func getArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article types.Article

		file, err := os.Open("./storage/" + r.PathValue("id") + ".json")
		if err != nil {
			log.Println(err.Error())
		}

		data, err := io.ReadAll(file)
		if err != nil {
			log.Println(err.Error())
		}

		json.Unmarshal(data, &article)

		defer file.Close()

		componant := templates.Article(article)
		componant.Render(context.Background(), w)
	}
}

func getAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article types.Article
		var articles []types.Article
		var file *os.File

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

		defer file.Close()

		componant := templates.Admin(articles)
		componant.Render(context.Background(), w)
	}
}

func getEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article types.Article

		file, err := os.Open("./storage/" + r.PathValue("id") + ".json")
		if err != nil {
			log.Println(err.Error())
		}

		data, err := io.ReadAll(file)
		if err != nil {
			log.Println(err.Error())
		}

		json.Unmarshal(data, &article)

		defer file.Close()

		componant := templates.Edit(article)
		componant.Render(context.Background(), w)
	}
}

func getDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := "./storage/" + r.PathValue("id") + ".json"

		err := os.Remove(fileName)
		if err != nil {
			log.Println(err.Error())
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
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

		defer file.Close()

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func postEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article types.Article

		r.ParseForm()
		article.Title = r.FormValue("title")
		article.Date = r.FormValue("publish_date")
		article.Content = r.FormValue("content")
		article.ID, _ = strconv.Atoi(r.PathValue("id"))

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

		defer file.Close()

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
