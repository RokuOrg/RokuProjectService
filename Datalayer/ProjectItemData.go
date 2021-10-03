package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func AddProjectItem(Item Models.ProjectItem, ListId string) Models.Message {

	Connect()
	defer db.Close()

	result, err := db.Query("INSERT INTO ProjectItem (Title, Description, ProjectListId, Position) VALUES (?, ?, ?, ?)", Item.Title, Item.Description, ListId, Item.Position)

	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	return Models.Message{
		Success: true,
		Value:   "Added projectList",
	}
}

func GetProjectItemsFromProjectList(Id string) Models.ProjectItems {

	Connect()
	defer db.Close()

	var projectItems Models.ProjectItems
	result, err := db.Query("SELECT * FROM ProjectItem WHERE ProjectListId = ?", Id)

	if err != nil {
		log.Fatal(err)

	}

	for result.Next() {
		var item Models.ProjectItem
		var projectListId string
		err := result.Scan(&item.Title, &item.Description, &projectListId, &item.Position)
		if err != nil {
			log.Fatal(err)
		}
		projectItems = append(projectItems, item)
	}

	return projectItems
}
