package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func (d *DatabaseLayer) AddProjectItem(Item Models.ProjectItemDAO, ListId string) error {

	defer d.DB.Close()

	result, err := d.DB.Query("INSERT INTO ProjectItem (Title, Description, ProjectListId, Position) VALUES (?, ?, ?, ?)", Item.Title, Item.Description, ListId, Item.Position)

	if err != nil {
		return err
	}
	defer result.Close()

	return nil
}

func (d *DatabaseLayer) GetProjectItemsFromProjectList(Id string) []Models.ProjectItemDAO {

	defer d.DB.Close()

	var projectItems []Models.ProjectItemDAO
	result, err := d.DB.Query("SELECT * FROM ProjectItem WHERE ProjectListId = ?", Id)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var item Models.ProjectItemDAO
		var projectListId string
		err := result.Scan(&item.Title, &item.Description, &projectListId, &item.Position)
		if err != nil {
			log.Fatal(err)
		}
		projectItems = append(projectItems, item)
	}

	return projectItems
}

func (d *DatabaseLayer) GetProjectItem(id string) Models.ProjectItemDAO {

	defer d.DB.Close()

	var item Models.ProjectItemDAO
	result, err := d.DB.Query("SELECT * FROM ProjectItem WHERE Id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var projlist string
		err := result.Scan(&item.Id, &item.Title, &item.Description, projlist, &item.Position)
		if err != nil {
			log.Fatal(err)
		}
	}
	return item
}
