package issue

import (
	"ghub/pkg/shared"
	"time"
)

type RequestedReviewers struct {
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

type License struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Spdx_id  string `json:"spdx_id"`
	Node_id  string `json:"node_id"`
	Html_url string `json:"html_url"`
}

type PrRepoData struct {
	Id                     int       `json:"id"`
	Node_id                string    `json:"node_id"`
	Name                   string    `json:"name"`
	Full_name              string    `json:"full_name"`
	Private                bool      `json:"private"`
	Html_url               string    `json:"html_url"`
	Description            string    `json:"description"`
	Fork                   bool      `json:"fork"`
	Url                    string    `json:"url"`
	Archive_url            string    `json:"archive_url"`
	Assignees_url          string    `json:"assignees_url"`
	Blobs_url              string    `json:"blobs_url"`
	Branches_url           string    `json:"branches_url"`
	Collaborators_url      string    `json:"collaborators_url"`
	Comments_url           string    `json:"comments_url"`
	Commits_url            string    `json:"commits_url"`
	Compare_url            string    `json:"compare_url"`
	Contents_url           string    `json:"contents_url"`
	Contributors_url       string    `json:"contributors_url"`
	Deployments_url        string    `json:"deployments_url"`
	Downloads_url          string    `json:"downloads_url"`
	Events_url             string    `json:"events_url"`
	Forks_url              string    `json:"forks_url"`
	Git_commits_url        string    `json:"git_commits_url"`
	Git_refs_url           string    `json:"git_refs_url"`
	Git_tags_url           string    `json:"git_tags_url"`
	Git_url                string    `json:"git_url"`
	Issue_comment_url      string    `json:"issue_comment_url"`
	Issue_events_url       string    `json:"issue_events_url"`
	Issues_url             string    `json:"issues_url"`
	Keys_url               string    `json:"keys_url"`
	Labels_url             string    `json:"labels_url"`
	Languages_url          string    `json:"languages_url"`
	Merges_url             string    `json:"merges_url"`
	Milestones_url         string    `json:"milestones_url"`
	Notifications_url      string    `json:"notifications_url"`
	Pulls_url              string    `json:"pulls_url"`
	Releases_url           string    `json:"releases_url"`
	Ssh_url                string    `json:"ssh_url"`
	Stargazers_url         string    `json:"stargazers_url"`
	Statuses_url           string    `json:"statuses_url"`
	Subscribers_url        string    `json:"subscribers_url"`
	Subscription_url       string    `json:"subscription_url"`
	Tags_url               string    `json:"tags_url"`
	Teams_url              string    `json:"teams_url"`
	Trees_url              string    `json:"trees_url"`
	Clone_url              string    `json:"clone_url"`
	Mirror_url             string    `json:"mirror_url"`
	Hooks_url              string    `json:"hooks_url"`
	Svn_url                string    `json:"svn_url"`
	Homepage               string    `json:"homepage"`
	Language               string    `json:"language"`
	Forks_count            int       `json:"forks_count"`
	Stargazers_count       int       `json:"stargazers_count"`
	Watchers_count         int       `json:"watchers_count"`
	Size                   int       `json:"size"`
	Default_branch         string    `json:"default_branch"`
	Open_issues_count      int       `json:"open_issues_count"`
	Is_template            bool      `json:"is_template"`
	Topics                 []string  `json:"topics"`
	Has_issues             bool      `json:"has_issues"`
	Has_projects           bool      `json:"has_projects"`
	Has_wiki               bool      `json:"has_wiki"`
	Has_pages              bool      `json:"has_pages"`
	Has_downloads          bool      `json:"has_downloads"`
	Archived               bool      `json:"archived"`
	Disabled               bool      `json:"disabled"`
	Visibility             string    `json:"visibility"`
	Pushed_at              time.Time `json:"pushed_at"`
	Created_at             time.Time `json:"created_at"`
	Updated_at             time.Time `json:"updated_at"`
	shared.Owner           `json:"owner"`
	shared.Permissions     `json:"permissions"`
	Allow_rebase_merge     bool   `json:"allow_rebase_merge"`
	Template_repository    string `json:"template_repository"`
	Temp_clone_token       string `json:"temp_clone_token"`
	Allow_squash_merge     bool   `json:"allow_squash_merge"`
	Allow_auto_merge       bool   `json:"allow_auto_merge"`
	Delete_branch_on_merge bool   `json:"delete_branch_on_merge"`
	Allow_merge_commit     bool   `json:"allow_merge_commit"`
	Subscribers_count      int    `json:"subscribers_count"`
	Network_count          int    `json:"network_count"`
	License                `json:"license"`
	Forks                  int `json:"forks"`
	Open_issues            int `json:"open_issues"`
	Watchers               int `json:"watchers"`
}

type Head struct {
	Label       string `json:"label"`
	Ref         string `json:"ref"`
	Sha         string `json:"sha"`
	shared.User `json:"user"`
	Repo        PrRepoData `json:"repo"`
}

type RequestedTeams struct {
	Id               int    `json:"id"`
	Node_id          string `json:"node_id"`
	Url              string `json:"url"`
	Html_url         string `json:"html_url"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Description      string `json:"description"`
	Privacy          string `json:"privacy"`
	Permission       string `json:"permission"`
	Members_url      string `json:"members_url"`
	Repositories_url string `json:"repositories_url"`
	Parent           string `json:"parent"`
}

type LinkHref struct {
	Href string `json:"href"`
}

type Links struct {
	Self            LinkHref `json:"self"`
	Html            LinkHref `json:"html"`
	Issue           LinkHref `json:"issue"`
	Comments        LinkHref `json:"comments"`
	Review_comments LinkHref `json:"review_comments"`
	Review_comment  LinkHref `json:"review_comment"`
	Commits         LinkHref `json:"commits"`
	Statuses        LinkHref `json:"statuses"`
}

type MergedBy struct {
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

type PullRequest struct {
	Url                   string `json:"url"`
	Id                    int    `json:"id"`
	Node_id               string `json:"Node_id"`
	Html_url              string `json:"html_url"`
	Diff_url              string `json:"diff_url"`
	Patch_url             string `json:"patch_url"`
	Issue_url             string `json:"issue_url"`
	Commits_url           string `json:"commits_url"`
	Review_comments_url   string `json:"review_comments_url"`
	Review_comment_url    string `json:"review_comment_url"`
	Comments_url          string `json:"comments_url"`
	Statuses_url          string `json:"statuses_url"`
	Number                int    `json:"number"`
	State                 string `json:"state"`
	Locked                bool   `json:"locked"`
	Title                 string `json:"title"`
	shared.User           `json:"user"`
	Body                  string          `json:"body"`
	Labels                []shared.Labels `json:"labels"`
	shared.Milestone      `json:"milestone"`
	Active_lock_reason    string `json:"active_lock_reason"`
	Created_at            string `json:"created_at"`
	Updated_at            string `json:"updated_at"`
	Closed_at             string `json:"closed_at"`
	Merged_at             string `json:"merged_at"`
	Merge_commit_sha      string `json:"merge_commit_sha"`
	shared.Assignee       `json:"assignee"`
	Assignees             []shared.Assignee    `json:"assignees"`
	Requested_reviewers   []RequestedReviewers `json:"requested_reviewers"`
	RequestedTeams        []RequestedTeams     `json:"requested_teams"`
	Head                  `json:"head"`
	Base                  Head `json:"base"`
	Links                 `json:"_links"`
	Author_association    string `json:"author_association"`
	Auto_merge            string `json:"auto_merge"`
	Draft                 bool   `json:"draft"`
	Merged                bool   `json:"merged"`
	Mergeable             bool   `json:"mergeable"`
	Rebaseable            bool   `json:"rebaseable"`
	Mergeable_state       string `json:"mergeable_state"`
	Comments              int    `json:"comments"`
	Review_comments       int    `json:"review_comments"`
	Maintainer_can_modify bool   `json:"maintainer_can_modify"`
	Commits               int    `json:"commits"`
	Additions             int    `json:"additions"`
	Deletions             int    `json:"deletions"`
	Changed_files         int    `json:"changed_files"`
	MergedBy              `json:"merged_by"`
}
