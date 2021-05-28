package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	log.Println("Starting API")
	r := mux.NewRouter()
	r.HandleFunc("/users/{user}/popular-repository", getPopularRepo).Methods("GET")
	r.HandleFunc("/users/{user}/repositories/{repository}/popular-issue", getMostPopularIssue).Methods("GET")
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

func buildUrl(urlSnippet string, complement ...string) (*url.URL, error) {
	var genericStrings []interface{} = make([]interface{}, len(complement))
	for i, s := range complement {
		genericStrings[i] = s
	}
	return url.Parse(fmt.Sprintf(urlSnippet, genericStrings...))
}

func fetch(url *url.URL) (*http.Response, error) {
	return http.Get(url.String())
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
