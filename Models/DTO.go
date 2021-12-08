package Models

type ProjectDTO struct {
	Name     string `json:"name"`
	Template bool   `json:"template"`
}

type ProjectListDTO struct {
	Position int    `json:"position"`
	Name     string `json:"name"`
}

type ProjectItemDTO struct {
	Title    string `json:"title"`
	Position int    `json:"position"`
}
