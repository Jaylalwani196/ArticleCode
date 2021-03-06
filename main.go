package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string //'json:"Id"'
	Title   string //'json:"Title"'
	Desc    string //'json:"Desc"'
	Content string //'json:"Content"'
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All")
	json.NewEncoder(w).Encode(Articles)

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint hit: get an article")
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Endpoint hit : Create new Article")
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Endpoint hit : Delete an article")
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequests() {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	fmt.Println("Starting web based simple program! ")
	Articles = []Article{
		Article{Id: "1", Title: "Harvey Specter", Desc: "Pearson Hardman", Content: "Great lawyer"},
		Article{Id: "2", Title: "Michael Ross", Desc: "Pearson specter litt", Content: "Fraud lawyer"},
	}
	handleRequests()
}
