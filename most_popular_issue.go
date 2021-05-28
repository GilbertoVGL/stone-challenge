package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Login               string `json:"login"`
	Id                  int    `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"Type"`
	Site_admin          bool   `json:"site_admin"`
}

type Labels struct {
	Id          int    `json:"id"`
	Node_id     string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"Default"`
}

type Assignee struct {
	Login               string `json:"login"`
	Id                  int    `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"type"`
	Site_admin          bool   `json:"site_admin"`
}

type Milestone struct {
	Url           string    `json:"url"`
	Html_url      string    `json:"html_url"`
	Labels_url    string    `json:"labels_url"`
	Id            int       `json:"id"`
	Node_id       string    `json:"node_id"`
	Number        int       `json:"number"`
	State         string    `json:"state"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Open_issues   int       `json:"open_issues"`
	Closed_issues int       `json:"closed_issues"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	Closed_at     time.Time `json:"closed_at"`
	Due_on        time.Time `json:"due_on"`
	Creator       `json:"creator"`
}

type Creator struct {
	Login               string `json:"login"`
	Id                  int    `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"Type"`
	Site_admin          bool   `json:"site_admin"`
}

type PullRequest struct {
	Url       string `json:"url"`
	Html_url  string `json:"html_url"`
	Diff_url  string `json:"diff_url"`
	Patch_url string `json:"patch_url"`
}

type IssueData struct {
	Id                 int       `json:"id"`
	Node_id            string    `json:"node_id"`
	Url                string    `json:"url"`
	Repository_url     string    `json:"repository_url"`
	Labels_url         string    `json:"labels_url"`
	Comments_url       string    `json:"comments_url"`
	Events_url         string    `json:"events_url"`
	Html_url           string    `json:"html_url"`
	Number             int       `json:"number"`
	State              string    `json:"state"`
	Title              string    `json:"title"`
	Body               string    `json:"body"`
	Locked             bool      `json:"locked"`
	Active_lock_reason string    `json:"active_lock_reason"`
	Comments           int       `json:"comments"`
	Closed_at          time.Time `json:"closed_at"`
	Created_at         time.Time `json:"created_at"`
	Updated_at         time.Time `json:"updated_at"`
	Author_association string    `json:"author_association"`
	User               `json:"user"`
	Labels             []Labels `json:"labels"`
	Assignee           `json:"assignee"`
	Assignees          []Assignee `json:"assignees"`
	Milestone          `json:"milestone"`
	PullRequest        `json:"pullRequest"`
}

const mostPopularIssueURL = "https://api.github.com/repos/%s/%s/issues?state=open&sort=comments&direction=desc&per_page=1"

func getMostPopularIssue(w http.ResponseWriter, r *http.Request) {
	log.Println("called getPopularIssue")
	var repoIssues []IssueData
	vars := mux.Vars(r)
	user := vars["user"]
	repo := vars["repository"]

	log.Println("user: ", user)
	log.Println("repo: ", repo)

	parsedUrl, err := buildUrl(mostPopularIssueURL, user, repo)

	if err != nil {
		log.Fatalln("error: ", err)
		respondWithError(w, http.StatusBadRequest, "unable to build github URL")
	}

	body, err := fetch(parsedUrl)

	if err != nil {
		log.Fatalln("error: ", err)
		respondWithError(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	err = parsePopularIssueResponse(body, &repoIssues)

	if err != nil {
		log.Fatalln("error: ", err)
		respondWithError(w, http.StatusBadRequest, "unable to parse github response")
	}

	log.Printf("Found %d open issues on %s\n", len(repoIssues), repo)
	repoData := sortByComments(repoIssues)

	respondWithJSON(w, http.StatusOK, repoData)

}

func parsePopularIssueResponse(resp *http.Response, issuesData *[]IssueData) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, issuesData)
}

func sortByComments(reposData []IssueData) IssueData {
	var issueData IssueData = reposData[0]

	log.Printf("sorting by comments count\n")

	log.Printf("Selected:\tcomments: %v\t %v\n", issueData.Comments, issueData.Title)

	return issueData
}
