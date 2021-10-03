package Datalayer

import (
	"RokuProject-Back-End/Models"
	"log"
)

func AddProjectUser(projectUser Models.ProjectUser) Models.Message {
	Connect()
	defer db.Close()

	result, err := db.Query("INSERT INTO ProjectUser (ProjectId, UserId ) VALUES (?, ?)", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Models.Message{
		Success: true,
		Value:   "Added projectUser",
	}
}

func RemoveProjectUser(projectUser Models.ProjectUser) Models.Message {
	Connect()
	defer db.Close()

	result, err := db.Query("DELETE FROM ProjectUser (ProjectId, UserId ) VALUES (?, ?)", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Models.Message{
		Success: true,
		Value:   "Deleted projectUser",
	}
}

func GetProjectUser(projectUser Models.ProjectUser) Models.ProjectUser {
	Connect()
	defer db.Close()

	var User Models.ProjectUser

	result, err := db.Query("SELECT * FROM ProjectUser WHERE ProjectId = ? AND UserId = ?", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		err := result.Scan(&User.ProjectId, &User.UserId)
		if err != nil {
			log.Fatal(err)
		}
	}

	return User
}

func GetProjectUsersByUserId(UserId string) Models.ProjectUsers {

	Connect()
	defer db.Close()

	var Users Models.ProjectUsers

	result, err := db.Query("SELECT * FROM ProjectUser WHERE UserId = ?", UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var User Models.ProjectUser
		err := result.Scan(&User.ProjectId, &User.UserId)
		if err != nil {
			log.Fatal(err)
		}
		Users = append(Users, User)
	}

	return Users
}

func GetAllProjectsFromUser(UserId string) Models.ProjectUsers {

	Connect()
	defer db.Close()

	var projectUsers Models.ProjectUsers

	result, err := db.Query("SELECT * FROM ProjectUser WHERE UserId = ?", UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var user Models.ProjectUser

		err := result.Scan(&user.ProjectId, &user.ProjectId)
		if err != nil {
			log.Fatal(err)
		}
		projectUsers = append(projectUsers, user)
	}

	return projectUsers
}
