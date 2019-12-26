package model

type Repo struct {
	Name       *string `json:"name"`
	Fullname   *string `json:"full_name"`
	Url        *string `json:"html_url"`
	Stars      *int    `json:"stargazers_count"`
	ForksCount *int    `json:"forks_count"`
}
