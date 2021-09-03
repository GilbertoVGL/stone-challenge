package listing

import "time"

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
