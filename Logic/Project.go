package Logic

import (
	"database/sql"
	"errors"
)

type Project struct {
	Id      string       `json:"id"`
	OwnerId string       `json:"userId"`
	Name    string       `json:"name"`
	Project ProjectLists `json:"project"`
}

func (p *Project) updateProject(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Project) addProjectList(db *sql.DB) error {
	return errors.New("Not implemented")
}
