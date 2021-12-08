package Logic

type ProjectItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}
