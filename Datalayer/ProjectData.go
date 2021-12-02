package Datalayer

import (
	"RokuProject-Back-End/Logic"
	"log"
)

func AddProject(project Logic.Project) Logic.Message {

	Connect()
	result, err := db.Query("INSERT INTO Project (Id , Name, OwnerId) VALUES (?, ?, ?)", project.Id, project.Name, project.OwnerId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()
	defer db.Close()
	return Logic.Message{
		Success: true,
		Value:   "Added project",
	}
}

func GetProjectFromId(Id string) Logic.Project {
	Connect()
	defer db.Close()
	var project Logic.Project
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

func GetAllProjects() Logic.Projects {
	Connect()
	defer db.Close()

	var projects Logic.Projects

	result, err := db.Query("SELECT * FROM Project")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var project Logic.Project

		err := result.Scan(&project.Id, &project.Name, &project.OwnerId)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	return projects
}
