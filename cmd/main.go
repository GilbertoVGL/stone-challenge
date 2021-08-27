package main

import (
	"ghub-api/pkg/listing"
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
	r.HandleFunc("/users/{user}/popular-repository", listing.GetPopularRepo).Methods("GET")
	r.HandleFunc("/users/{user}/repositories/{repository}/popular-issue", listing.GetMostPopularIssue).Methods("GET")
	walk(r)
	log.Fatal(http.ListenAndServe(":8080", r))
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
