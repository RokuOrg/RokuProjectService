package Logic

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Models"
	"database/sql"
	"errors"
	"github.com/segmentio/ksuid"
)

type ProjectCollection struct {
	projects  []Project
	DataLayer *Datalayer.DatabaseLayer
	RandomId  func() ksuid.KSUID
}

func (p *ProjectCollection) CreateProject(UserId string, ProjectName string, template bool) error {
	if p.DataLayer == nil {
		return errors.New("Collection not initialized")
	}

	project := Project{
		Name:    ProjectName,
		Id:      p.RandomId().String(),
		OwnerId: UserId,
		Project: nil,
	}

	if template {
		if err := project.CreateDefaultLists(); err != nil {
			return err
		}
	}

	if err := p.DataLayer.AddProject(Models.ProjectDAO{Name: project.Name, Id: project.Id, OwnerId: project.OwnerId}); err != nil {
		return err
	}

	//.DataLayer.AddProjectUser(ProjectUser{ProjectId: project.Id, UserId: UserId})

	return nil
}

func (p *ProjectCollection) GetProjects() ([]Project, error) {
	if p.DataLayer == nil {
		return []Project{}, errors.New("Collection not initialized")
	}

	projs := p.DataLayer.GetAllProjects()
	var projects []Project

	for i := range projs {
		pr, err := p.GetProject(projs[i].Id)
		if err != nil {
			return nil, err
		}

		projects = append(projects, pr)
	}

	return projects, nil
}

func (p *ProjectCollection) ConvertProjectToDAO(project Project) Models.ProjectDAO {
	return Models.ProjectDAO{
		Name:    project.Name,
		Id:      project.Id,
		OwnerId: project.OwnerId,
	}
}

func (p *ProjectCollection) ConvertDTOtoProject(dto Models.ProjectDTO) Project {
	return Project{
		Id:   p.RandomId().String(),
		Name: dto.Name,
	}
}

func (p *ProjectCollection) GetProject(id string) (Project, error) {
	if p.DataLayer == nil {
		return Project{}, errors.New("Collection not initialized")
	}

	return Project{}, errors.New("Not implemented")
}

func (p *ProjectCollection) DeleteProject() error {
	if p.DataLayer == nil {
		return errors.New("Collection not initialized")
	}

	return errors.New("Not implemented")
}

func (p *ProjectCollection) testMethod() {

}

func BuildProjectCollection(db *sql.DB, randomId func() ksuid.KSUID) *ProjectCollection {
	Dl := Datalayer.CreateDatabaseLayer(db)
	p := ProjectCollection{DataLayer: Dl, RandomId: randomId}

	return &p
}
