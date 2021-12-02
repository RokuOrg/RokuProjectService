package Datalayer

import (
	"RokuProject-Back-End/Logic"
	"fmt"
	"log"
)

func AddProjectUser(projectUser Logic.ProjectUser) Logic.Message {
	Connect()
	defer db.Close()

	result, err := db.Query("INSERT INTO ProjectUser (ProjectId, UserId ) VALUES (?, ?)", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Logic.Message{
		Success: true,
		Value:   "Added projectUser",
	}
}

func RemoveProjectUser(projectUser Logic.ProjectUser) Logic.Message {
	Connect()
	defer db.Close()

	result, err := db.Query("DELETE FROM ProjectUser (ProjectId, UserId ) VALUES (?, ?)", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	return Logic.Message{
		Success: true,
		Value:   "Deleted projectUser",
	}
}

func GetProjectUser(projectUser Logic.ProjectUser) Logic.ProjectUser {
	Connect()
	defer db.Close()

	var User Logic.ProjectUser

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

func GetProjectUsersByUserId(UserId string) Logic.ProjectUsers {

	Connect()
	defer db.Close()

	var Users Logic.ProjectUsers

	result, err := db.Query("SELECT * FROM ProjectUser WHERE UserId = ?", UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var User Logic.ProjectUser
		err := result.Scan(&User.ProjectId, &User.UserId)
		if err != nil {
			log.Fatal(err)
		}
		Users = append(Users, User)
	}

	return Users
}

func GetAllProjectsFromUser(UserId string) Logic.ProjectUsers {

	Connect()
	defer db.Close()

	var projectUsers Logic.ProjectUsers

	result, err := db.Query("SELECT * FROM ProjectUser WHERE UserId = ?", UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var user Logic.ProjectUser

		err := result.Scan(&user.ProjectId, &user.UserId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		projectUsers = append(projectUsers, user)
	}

	return projectUsers
}
