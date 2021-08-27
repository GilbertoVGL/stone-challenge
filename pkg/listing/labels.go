package listing

type Labels struct {
	Id          int    `json:"id"`
	Node_id     string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"Default"`
}
