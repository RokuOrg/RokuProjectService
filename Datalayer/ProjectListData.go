package Datalayer

import (
	"RokuProject-Back-End/Logic"
	"log"
)

func AddProjectList(list Logic.ProjectList, ProjectId string) Logic.Message {

	Connect()
	defer db.Close()

	result, err := db.Query("INSERT INTO ProjectList (Id , Name, ProjectId, Position ) VALUES (?, ?, ?, ?)", list.Id, list.Name, ProjectId, list.Position)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Logic.Message{
		Success: true,
		Value:   "Added projectList",
	}
}

func GetProjectListsFromProjectId(Id string) Logic.ProjectLists {
	Connect()
	defer db.Close()
	var projectLists Logic.ProjectLists
	result, err := db.Query("SELECT * FROM ProjectList WHERE ProjectId = ?", Id)

	if err != nil {
		log.Fatal(err)
		return Logic.ProjectLists{}
	}

	for result.Next() {
		var list Logic.ProjectList
		var projectId string
		err := result.Scan(&list.Id, &list.Name, &projectId, &list.Position)
		if err != nil {
			log.Fatal(err)
		}
		projectLists = append(projectLists, list)
	}

	return projectLists
}

func RemoveProjectList(ProjectId string) Logic.Message {

	Connect()
	defer db.Close()

	result, err := db.Query("DELETE FROM ProjectList WHERE ProjectId = ?", ProjectId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Logic.Message{
		Success: true,
		Value:   "Removed ProjectList",
	}
}
