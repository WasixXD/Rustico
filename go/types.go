package main

type ProjectRequest struct {
	Project_name string `json:"project_name"`
	Description  string `json:"description"`
}

type ProjectResponse struct {
	Id        int     `json:"id"`
	Name      string  `json:"project_name"`
	Desc      string  `json:"description"`
	Created   string  `json:"created_at"`
	GithubUrl *string `json:"github_url"`
}
