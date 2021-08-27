package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func BuildUrl(urlSnippet string, complement ...string) (*url.URL, error) {
	var genericStrings []interface{} = make([]interface{}, len(complement))
	for i, s := range complement {
		genericStrings[i] = s
	}
	return url.Parse(fmt.Sprintf(urlSnippet, genericStrings...))
}

func Fetch(url *url.URL) (*http.Response, error) {
	return http.Get(url.String())
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
