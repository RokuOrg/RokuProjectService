package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func (d *DatabaseLayer) AddProjectList(list Models.ProjectListDAO, ProjectId string) error {

	defer d.DB.Close()

	result, err := d.DB.Query("INSERT INTO ProjectList (Id , Name, ProjectId, Position ) VALUES (?, ?, ?, ?)", list.Id, list.Name, ProjectId, list.Position)

	if err != nil {
		return err
	}

	defer result.Close()

	return nil
}

func (d *DatabaseLayer) GetProjectListsFromProjectId(Id string) []Models.ProjectListDAO {

	defer d.DB.Close()
	var projectLists []Models.ProjectListDAO
	result, err := d.DB.Query("SELECT * FROM ProjectList WHERE ProjectId = ?", Id)

	if err != nil {
		log.Fatal(err)
		return []Models.ProjectListDAO{}
	}

	for result.Next() {
		var list Models.ProjectListDAO
		var projectId string
		err := result.Scan(&list.Id, &list.Name, &projectId, &list.Position)
		if err != nil {
			log.Fatal(err)
		}
		projectLists = append(projectLists, list)
	}

	return projectLists
}

func (d *DatabaseLayer) RemoveProjectList(ProjectId string) error {

	defer d.DB.Close()

	result, err := d.DB.Query("DELETE FROM ProjectList WHERE ProjectId = ?", ProjectId)

	if err != nil {
		return err
	}

	defer result.Close()

	return nil
}
