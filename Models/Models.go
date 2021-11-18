package Models

type ProjectUsers []ProjectUser

type ProjectUser struct {
	ProjectId string `json:"projectId"`
	UserId    string `json:"userId"`
}

type Projects []Project

type Project struct {
	Id      string       `json:"id"`
	OwnerId string       `json:"userId"`
	Name    string       `json:"name"`
	Project ProjectLists `json:"project"`
}

type ProjectLists []ProjectList

type ProjectList struct {
	Id       string       `json:"id"`
	Position int          `json:"position"`
	Name     string       `json:"name"`
	Items    ProjectItems `json:"items"`
}

type ProjectItems []ProjectItem

type ProjectItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}

type CreateProject struct {
	Name     string `json:"name"`
	Template bool   `json:"template"`
}

type Message struct {
	Success bool   `json:"success"`
	Value   string `json:"value"`
}

type CreateProjectList struct {
	Name string `json:"name"`
}

type CreateProjectItem struct {
	Title    string `json:"title"`
	Position int    `json:"position"`
}
