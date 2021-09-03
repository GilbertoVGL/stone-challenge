package listing

import "time"

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
	PullRequest        `json:"pull_request"`
}
