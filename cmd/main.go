package main

import (
	"ghub/pkg/issue"
	"ghub/pkg/pullrequest"
	"ghub/pkg/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	log.Println("Starting API")
	r := mux.NewRouter()
	r.HandleFunc("/users/{user}/popular-repository", repository.GetPopular).Methods("GET")
	r.HandleFunc("/users/{user}/repositories/{repository}/popular-issue", issue.GetPopular).Methods("GET")
	r.HandleFunc("/users/{user}/repositories/{repository}/open-pull-requests", pullrequest.GetOpen).Methods("GET")
	walk(r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func walk(router *mux.Router) {
	log.Printf("Available routes: \n")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		log.Printf("\t %v - %v - %v - %v \n", tpl, err1, met, err2)
		return nil
	})
}
