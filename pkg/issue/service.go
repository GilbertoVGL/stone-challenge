package issue

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

const URL = "https://api.github.com/repos/%s/%s/issues?state=open&sort=comments&direction=desc&per_page=1"

var relRegex = regexp.MustCompile(`&page=(\d+)>; rel="last"`)
var lastPageRegex = regexp.MustCompile(`\d+`)

func GetPopular(w http.ResponseWriter, r *http.Request) {
	log.Println("called get popular issue")
	var repoIssues []IssueData
	vars := mux.Vars(r)
	user := vars["user"]
	repo := vars["repository"]

	log.Println("user: ", user)
	log.Println("repo: ", repo)

	parsedUrl, err := rest.NewURL(URL, user, repo)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to build github URL")
	}

	body, err := rest.Get(parsedUrl)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	err = parse(body, &repoIssues)

	if err != nil {
		log.Fatalln("error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to parse github response")
	}

	log.Printf("Found %d open issues on %s\n", len(repoIssues), repo)
	repoData := sort(repoIssues)

	rest.JSONresp(w, http.StatusOK, repoData)
}

func parse(resp *http.Response, issuesData *[]IssueData) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, issuesData)
}

func sort(reposData []IssueData) IssueData {
	var issueData IssueData = reposData[0]

	log.Printf("sorting by comments count\n")

	log.Printf("Selected:\tcomments: %v\t %v\n", issueData.Comments, issueData.Title)

	return issueData
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
