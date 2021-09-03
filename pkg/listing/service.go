package listing

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"

	"ghub-api/pkg/rest"
)

const mostPopularIssueURL = "https://api.github.com/repos/%s/%s/issues?state=open&sort=comments&direction=desc&per_page=1"
const starRepoURL = "https://api.github.com/users/%s/repos?sort=stargazers_count&direction=desc&per_page=1"
const openPRURL = "https://api.github.com/repos/%s/%s/pulls?state=open&sort=popularity&direction=asc&per_page=2&page=1"

var relRegex = regexp.MustCompile(`&page=(\d+)>; rel="last"`)
var lastPageRegex = regexp.MustCompile(`\d+`)

func GetMostPopularIssue(w http.ResponseWriter, r *http.Request) {
	log.Println("called get popular issue")
	var repoIssues []IssueData
	vars := mux.Vars(r)
	user := vars["user"]
	repo := vars["repository"]

	log.Println("user: ", user)
	log.Println("repo: ", repo)

	parsedUrl, err := rest.BuildUrl(mostPopularIssueURL, user, repo)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to build github URL")
	}

	body, err := rest.Fetch(parsedUrl)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	err = parsePopularIssueResponse(body, &repoIssues)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to parse github response")
	}

	log.Printf("Found %d open issues on %s\n", len(repoIssues), repo)
	repoData := sortByComments(repoIssues)

	rest.RespondWithJSON(w, http.StatusOK, repoData)

}

func parsePopularIssueResponse(resp *http.Response, issuesData *[]IssueData) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, issuesData)
}

func GetPopularRepo(w http.ResponseWriter, r *http.Request) {
	log.Println("called get popular repo")

	var userRepositories []RepoData
	vars := mux.Vars(r)
	user := vars["user"]

	log.Println("user: ", user)

	parsedUrl, err := rest.BuildUrl(starRepoURL, user)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to build github URL")
	}

	body, err := rest.Fetch(parsedUrl)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	err = parsePopularRepoResponse(body, &userRepositories)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to parse github response")
	}

	log.Printf("Found %d repositories of %s\n", len(userRepositories), user)
	repoData := sortRepoByStars(userRepositories)

	rest.RespondWithJSON(w, http.StatusOK, repoData)
}

func parsePopularRepoResponse(resp *http.Response, reposData *[]RepoData) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, reposData)
}

func GetOpenPR(w http.ResponseWriter, r *http.Request) {
	log.Println("called get open pull request")

	var repoPRs []PullRequest
	vars := mux.Vars(r)
	user := vars["user"]
	repo := vars["repository"]

	log.Println("user: ", user)
	log.Println("repo: ", repo)

	parsedUrl, err := rest.BuildUrl(openPRURL, user, repo)
	log.Println("parsedUrl: ", parsedUrl)

	if err != nil {
		log.Fatalln("BuildUrl error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to build github URL")
	}

	resp, err := rest.Fetch(parsedUrl)

	if err != nil {
		log.Fatalln("Fetch error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	pagesToGo, err := getNumOfPages(resp)

	log.Println("pagesToGo: ", pagesToGo)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to parse last page")
	}

	err = parseOpenPRResponse(resp, &repoPRs)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.RespondWithError(w, http.StatusBadRequest, "unable to parse github response")
	}

	rest.RespondWithJSON(w, http.StatusOK, repoPRs)
}

func parseOpenPRResponse(resp *http.Response, reposData *[]PullRequest) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, reposData)
}

func fetchPaginated() {

}

func sortRepoByStars(reposData []RepoData) RepoData {
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

func sortByComments(reposData []IssueData) IssueData {
	var issueData IssueData = reposData[0]

	log.Printf("sorting by comments count\n")

	log.Printf("Selected:\tcomments: %v\t %v\n", issueData.Comments, issueData.Title)

	return issueData
}

func getNumOfPages(r *http.Response) (int, error) {
	link := r.Header.Get("Link")

	rel := relRegex.FindString(link)
	lastPageStr := lastPageRegex.FindString(rel)

	lastPage, err := strconv.Atoi(lastPageStr)

	if err != nil {
		return 0, err
	}

	log.Println("lastPage: ", lastPage)

	return lastPage, nil
}
