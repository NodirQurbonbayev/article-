package handlers

import (
	"encoding/json"
	"httpserver/models"
	"net/http"
)

func HandleArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createArticle(w, r)
	case http.MethodGet:
		getArticles(w, r)
	case http.MethodPut:
		updateArticle(w, r)
	case http.MethodDelete:
		deleteArticle(w, r)

	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var newarticle models.Article
	err := json.NewDecoder(r.Body).Decode(&newarticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	newarticle.ID = len(models.Articles) + 1
	models.Articles = append(models.Articles, newarticle)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Articles)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(models.Articles)	
}

func updateArticle(w http.ResponseWriter, r *http.Request) {

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

}
