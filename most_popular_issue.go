package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const mostPopularIssueURL = "https://api.github.com/repos/%s/%s/issues"

func getMostPopularIssue(w http.ResponseWriter, r *http.Request) {
	log.Println("called getPopularIssue")
	vars := mux.Vars(r)
	user := vars["user"]
	repo := vars["repository"]

	log.Println("user: ", user)
	log.Println("repo: ", repo)

	url, err := buildUrl()

}
