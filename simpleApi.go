package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string
	Title   string
	Desc    string
	Content string
}

var Articles []Article

func homPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homPage")
	fmt.Fprintf(w, "hello from go")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return All Articles")
	json.NewEncoder(w).Encode(Articles)

}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return one Article")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}

}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return one Article")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode(article)
		}
	}

}
func addArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: add one Article")
	newArticleBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var newArticle Article
	json.Unmarshal(newArticleBody, &newArticle)
	Articles = append(Articles, newArticle)
	fmt.Println(newArticle)

}
func editArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: edit one Article")
	editArticleBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var editArticle Article
	json.Unmarshal(editArticleBody, &editArticle)

	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles[index] = editArticle
		}
	}
	fmt.Println("editArticle")

}
func handelRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homPage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", addArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", editArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8585", myRouter))
}

func main() {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description2", Content: "Article Content2"},
		{Id: "3", Title: "Hello 3", Desc: "Article Description3", Content: "Article Content3"},
	}
	handelRequest()
}
