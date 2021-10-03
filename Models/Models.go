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
	Position    int    `json:"position"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateProject struct {
	Name     string `json:"name"`
	Token    string `json:"Token"`
	Template bool   `json:"template"`
}

type Message struct {
	Success bool   `json:"success"`
	Value   string `json:"value"`
}

type CreateProjectList struct {
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
	Token     string `json:"token"`
}

type CreateProjectItem struct {
	ProjectId   string `json:"projectId"`
	ListId      string `json:"listId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Token       string `json:"token"`
}

type ResponseMessage struct {
	Succes bool   `json:"succes"`
	Object Object `json:"object"`
}

type Object struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Error    string `json:"error"`
	Id       int    `json:"id"`
	Message  string `json:"message"`
}
