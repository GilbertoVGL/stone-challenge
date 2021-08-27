package listing

import "time"

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
