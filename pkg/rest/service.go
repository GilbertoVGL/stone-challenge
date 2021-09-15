package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func NewURL(urlSnippet string, complement ...interface{}) (*url.URL, error) {
	var genericStrings = make([]interface{}, len(complement))
	for i, s := range complement {
		switch v := s.(type) {
		case int:
			genericStrings[i] = strconv.FormatInt(int64(v), 10)
		case string:
			genericStrings[i] = v
		}
	}
	return url.Parse(fmt.Sprintf(urlSnippet, genericStrings...))
}

func Get(url *url.URL) (*http.Response, error) {
	return http.Get(url.String())
}

func ErrorResp(w http.ResponseWriter, code int, message string) {
	JSONresp(w, code, map[string]string{"error": message})
}

func JSONresp(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
