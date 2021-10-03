package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func AddProjectList(list Models.ProjectList, ProjectId string) Models.Message {

	Connect()
	defer db.Close()

	result, err := db.Query("INSERT INTO ProjectList (Id , Name, ProjectId, Position ) VALUES (?, ?, ?, ?)", list.Id, list.Name, ProjectId, list.Position)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Models.Message{
		Success: true,
		Value:   "Added projectList",
	}
}

func GetProjectListsFromProjectId(Id string) Models.ProjectLists {
	Connect()
	defer db.Close()
	var projectLists Models.ProjectLists
	result, err := db.Query("SELECT * FROM ProjectList WHERE ProjectId = ?", Id)

	if err != nil {
		log.Fatal(err)
		return Models.ProjectLists{}
	}

	for result.Next() {
		var list Models.ProjectList
		var projectId string
		err := result.Scan(&list.Id, &list.Name, &projectId, &list.Position)
		if err != nil {
			log.Fatal(err)
		}
		projectLists = append(projectLists, list)
	}

	return projectLists
}

func RemoveProjectList(ProjectId string) Models.Message {

	Connect()
	defer db.Close()

	result, err := db.Query("DELETE FROM ProjectList WHERE ProjectId = ?", ProjectId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Models.Message{
		Success: true,
		Value:   "Removed ProjectList",
	}
}
