package github

type Repo struct {
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	PushedAt  string `json:"pushed_at"`
	FullName  string `json:"full_name"`
	Url       string `json:"html_url"`
	Stars     int64  `json:"stargazers_count"`
	Forks     int64  `json:"forks_count"`
	Watchers  int64  `json:"watchers_count"`
	Issues    int64  `json:"open_issues_count"`
}
