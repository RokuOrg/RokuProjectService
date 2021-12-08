package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func (d *DatabaseLayer) AddProject(project Models.ProjectDAO) error {
	result, err := d.DB.Query("INSERT INTO Project (Id , Name, OwnerId) VALUES (?, ?, ?)", project.Id, project.Name, project.OwnerId)

	if err != nil {
		return err
	}

	defer result.Close()
	defer d.DB.Close()
	return nil
}

func (d *DatabaseLayer) GetProjectFromId(Id string) Models.ProjectDAO {

	defer d.DB.Close()
	var project Models.ProjectDAO
	result, err := d.DB.Query("SELECT * FROM Project WHERE Id = ?", Id)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		err := result.Scan(&project.Id, &project.Name, &project.OwnerId)
		if err != nil {
			log.Fatal(err)
		}
	}

	return project
}

func (d *DatabaseLayer) GetAllProjects() []Models.ProjectDAO {

	defer d.DB.Close()

	var projects []Models.ProjectDAO

	result, err := d.DB.Query("SELECT * FROM Project")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var project Models.ProjectDAO

		err := result.Scan(&project.Id, &project.Name, &project.OwnerId)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	return projects
}
