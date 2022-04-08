package models

type Node struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ForksCount  int    `json:"forksCount"`
}

type Projects struct {
	Nodes []Node `json:"nodes"`
}

type GetProjects struct {
	Projects Projects `json:"projects"`
}

type ProjectsWithRepoCount struct {
	Names string `json:"project_names"`
	Count int    `json:"total_fork_count"`
}
