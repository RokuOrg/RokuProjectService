package Models

type ProjectDAO struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	OwnerId string `json:"ownerId"`
}

type ProjectUserDAO struct {
	ProjectId string `json:"projectId"`
	UserId    string `json:"userId"`
}

type ProjectListDAO struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ProjectId string `json:"projectId"`
	Position  string `json:"position"`
}

type ProjectItemDAO struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ProjectListID string `json:"projectListId"`
	Position      string `json:"position"`
}
