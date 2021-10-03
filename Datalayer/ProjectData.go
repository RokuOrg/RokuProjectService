package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func AddProject(project Models.Project) Models.Message {

	Connect()
	result, err := db.Query("INSERT INTO Project (Id , Name, OwnerId) VALUES (?, ?, ?)", project.Id, project.Name, project.OwnerId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()
	defer db.Close()
	return Models.Message{
		Success: true,
		Value:   "Added project",
	}
}

func GetProjectFromId(Id string) Models.Project {
	Connect()
	defer db.Close()
	var project Models.Project
	result, err := db.Query("SELECT * FROM Project WHERE Id = ?", Id)

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

func GetAllProjects() Models.Projects {
	Connect()
	defer db.Close()

	var projects Models.Projects

	result, err := db.Query("SELECT * FROM Project")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var project Models.Project

		err := result.Scan(&project.Id, &project.Name, &project.OwnerId)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	return projects
}
