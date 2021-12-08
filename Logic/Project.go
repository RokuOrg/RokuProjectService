package Logic

import (
	"errors"
)

type Project struct {
	Id      string        `json:"id"`
	OwnerId string        `json:"userId"`
	Name    string        `json:"name"`
	Project []ProjectList `json:"project"`
	Users   []string      `json:"users"`
}

func (p *Project) UpdateProject() error {
	return errors.New("Not implemented")
}

func (p *Project) AddProjectList() error {
	return errors.New("Not implemented")
}

func (p *Project) AddTemplate() error {
	return errors.New("Not implemented")
}
