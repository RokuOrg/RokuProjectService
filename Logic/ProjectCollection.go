package Logic

import (
	"database/sql"
	"errors"
)

type ProjectCollection struct {
	projects []Project
	DB       *sql.DB
}

func (p *ProjectCollection) GetProject(id string) (Project, error) {
	if p.DB == nil {
		return Project{}, errors.New("Collection not initialized")
	}

	return Project{}, errors.New("Not implemented")
}

func (p *ProjectCollection) DeleteProject() error {
	if p.DB == nil {
		return errors.New("Collection not initialized")
	}

	return errors.New("Not implemented")
}

func (p *ProjectCollection) CreateProject() error {
	if p.DB == nil {
		return errors.New("Collection not initialized")
	}

	return errors.New("Not implemented")
}

func (p *ProjectCollection) GetProjects() ([]Project, error) {
	if p.DB == nil {
		return []Project{}, errors.New("Collection not initialized")
	}

	return nil, errors.New("Not implemented")
}

func BuildProjectCollection(db *sql.DB) *ProjectCollection {
	p := ProjectCollection{DB: db}

	return &p
}
