package Logic

type ProjectList struct {
	Id       string        `json:"id"`
	Position int           `json:"position"`
	Name     string        `json:"name"`
	Items    []ProjectItem `json:"items"`
}

func (p *ProjectList) AddProjectItem() {

}
