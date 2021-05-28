package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Owner struct {
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

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

type RepoData struct {
	Id                int       `json:"id"`
	Node_id           string    `json:"node_id"`
	Name              string    `json:"name"`
	Full_name         string    `json:"full_name"`
	Private           bool      `json:"private"`
	Html_url          string    `json:"html_url"`
	Description       string    `json:"description"`
	Fork              bool      `json:"fork"`
	Url               string    `json:"url"`
	Archive_url       string    `json:"archive_url"`
	Assignees_url     string    `json:"assignees_url"`
	Blobs_url         string    `json:"blobs_url"`
	Branches_url      string    `json:"branches_url"`
	Collaborators_url string    `json:"collaborators_url"`
	Comments_url      string    `json:"comments_url"`
	Commits_url       string    `json:"commits_url"`
	Compare_url       string    `json:"compare_url"`
	Contents_url      string    `json:"contents_url"`
	Contributors_url  string    `json:"contributors_url"`
	Deployments_url   string    `json:"deployments_url"`
	Downloads_url     string    `json:"downloads_url"`
	Events_url        string    `json:"events_url"`
	Forks_url         string    `json:"forks_url"`
	Git_commits_url   string    `json:"git_commits_url"`
	Git_refs_url      string    `json:"git_refs_url"`
	Git_tags_url      string    `json:"git_tags_url"`
	Git_url           string    `json:"git_url"`
	Issue_comment_url string    `json:"issue_comment_url"`
	Issue_events_url  string    `json:"issue_events_url"`
	Issues_url        string    `json:"issues_url"`
	Keys_url          string    `json:"keys_url"`
	Labels_url        string    `json:"labels_url"`
	Languages_url     string    `json:"languages_url"`
	Merges_url        string    `json:"merges_url"`
	Milestones_url    string    `json:"milestones_url"`
	Notifications_url string    `json:"notifications_url"`
	Pulls_url         string    `json:"pulls_url"`
	Releases_url      string    `json:"releases_url"`
	Ssh_url           string    `json:"ssh_url"`
	Stargazers_url    string    `json:"stargazers_url"`
	Statuses_url      string    `json:"statuses_url"`
	Subscribers_url   string    `json:"subscribers_url"`
	Subscription_url  string    `json:"subscription_url"`
	Tags_url          string    `json:"tags_url"`
	Teams_url         string    `json:"teams_url"`
	Trees_url         string    `json:"trees_url"`
	Clone_url         string    `json:"clone_url"`
	Mirror_url        string    `json:"mirror_url"`
	Hooks_url         string    `json:"hooks_url"`
	Svn_url           string    `json:"svn_url"`
	Homepage          string    `json:"homepage"`
	Language          string    `json:"language"`
	Forks_count       int       `json:"forks_count"`
	Stargazers_count  int       `json:"stargazers_count"`
	Watchers_count    int       `json:"watchers_count"`
	Size              int       `json:"size"`
	Default_branch    string    `json:"default_branch"`
	Open_issues_count int       `json:"open_issues_count"`
	Is_template       bool      `json:"is_template"`
	Topics            []string  `jstopicson:""`
	Has_issues        bool      `json:"has_issues"`
	Has_projects      bool      `json:"has_projects"`
	Has_wiki          bool      `json:"has_wiki"`
	Has_pages         bool      `json:"has_pages"`
	Has_downloads     bool      `json:"has_downloads"`
	Archived          bool      `json:"archived"`
	Disabled          bool      `json:"disabled"`
	Visibility        string    `json:"visibility"`
	Pushed_at         time.Time `json:"pushed_at"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
	Owner             `json:"owner"`
	Permissions       `json:"permissions"`
}

const starRepoURL = "https://api.github.com/users/%s/repos?sort=stargazers_count&direction=desc&per_page=1"

func getPopularRepo(w http.ResponseWriter, r *http.Request) {
	log.Println("called getPopularRepo")

	var userRepositories []RepoData
	vars := mux.Vars(r)
	user := vars["user"]

	log.Println("user: ", user)

	parsedUrl, err := buildUrl(starRepoURL, user)

	if err != nil {
		log.Fatalln("error: ", err)
		respondWithError(w, http.StatusBadRequest, "unable to build github URL")
	}

	body, err := fetch(parsedUrl)

	if err != nil {
		log.Fatalln("error: ", err)
		respondWithError(w, http.StatusBadRequest, "unable to fetch user repositories")
	}

	err = parsePopularRepoResponse(body, &userRepositories)

	if err != nil {
		log.Fatalln("error: ", err)
		respondWithError(w, http.StatusBadRequest, "unable to parse github response")
	}

	log.Printf("Found %d repositories of %s\n", len(userRepositories), user)
	repoData := sortRepoByStars(userRepositories)

	respondWithJSON(w, http.StatusOK, repoData)
}

func parsePopularRepoResponse(resp *http.Response, reposData *[]RepoData) error {
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
