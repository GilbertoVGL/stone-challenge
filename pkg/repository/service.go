package repository

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"

	"ghub/pkg/rest"
)

const URL = "https://api.github.com/users/%s/repos?sort=stargazers_count&direction=desc&per_page=1"

var relRegex = regexp.MustCompile(`&page=(\d+)>; rel="last"`)
var lastPageRegex = regexp.MustCompile(`\d+`)

func GetPopular(w http.ResponseWriter, r *http.Request) {
	log.Println("called get popular repo")

	var userRepositories []RepoData
	vars := mux.Vars(r)
	user := vars["user"]

	log.Println("user: ", user)

	parsedUrl, err := rest.NewURL(URL, user)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to build github URL")
	}

	body, err := rest.Get(parsedUrl)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	err = parse(body, &userRepositories)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to parse github response")
	}

	log.Printf("Found %d repositories of %s\n", len(userRepositories), user)
	repoData := sort(userRepositories)

	rest.JSONresp(w, http.StatusOK, repoData)
}

func parse(resp *http.Response, reposData *[]RepoData) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, reposData)
}

func sort(reposData []RepoData) RepoData {
	var repoData RepoData

	log.Printf("sorting by stars count\n")
	for _, r := range reposData {
		log.Printf("\t%c %v\t %v\n", 9734, r.Stargazers_count, r.Name)

		if r.Stargazers_count > repoData.Stargazers_count {
			repoData = r
		}
	}

	log.Printf("Selected:\t%c %v\t %v\n", 9734, repoData.Stargazers_count, repoData.Name)

	return repoData
}

func lastPage(r *http.Response) (int, error) {
	link := r.Header.Get("Link")

	a := r.Header.Values("Link")
	log.Println("TESTE: ", a)

	rel := relRegex.FindString(link)
	lastPageStr := lastPageRegex.FindString(rel)

	if lastPageStr == "" {
		return 0, nil
	}

	lastPage, err := strconv.Atoi(lastPageStr)

	if err != nil {
		return 0, err
	}

	log.Println("lastPage: ", lastPage)

	return lastPage, nil
}
