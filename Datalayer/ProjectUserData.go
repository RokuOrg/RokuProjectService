package Datalayer

import (
	"RokuProject-Back-End/Models"
	"fmt"
	"log"
)

func (d *DatabaseLayer) AddProjectUser(projectUser Models.ProjectUserDAO) error {

	defer d.DB.Close()

	result, err := d.DB.Query("INSERT INTO ProjectUser (ProjectId, UserId ) VALUES (?, ?)", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		return err
	}

	defer result.Close()

	return nil
}

func (d *DatabaseLayer) RemoveProjectUser(projectUser Models.ProjectUserDAO) error {

	defer d.DB.Close()

	result, err := d.DB.Query("DELETE FROM ProjectUser WHERE ProjectId =? AND UserId = ? ", projectUser.ProjectId, projectUser.UserId)

	if err != nil {
		return err
	}

	defer result.Close()

	return nil
}

func (d *DatabaseLayer) GetProjectUser(projectUser Models.ProjectUserDAO) Models.ProjectUserDAO {

	defer d.DB.Close()

	var User Models.ProjectUserDAO

	result, err := d.DB.Query("SELECT * FROM ProjectUser WHERE ProjectId = ? AND UserId = ?", projectUser.ProjectId, projectUser.UserId)

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

func (d *DatabaseLayer) GetProjectUsersByUserId(UserId string) []Models.ProjectUserDAO {

	defer d.DB.Close()

	var Users []Models.ProjectUserDAO

	result, err := d.DB.Query("SELECT * FROM ProjectUser WHERE UserId = ?", UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var User Models.ProjectUserDAO
		err := result.Scan(&User.ProjectId, &User.UserId)
		if err != nil {
			log.Fatal(err)
		}
		Users = append(Users, User)
	}

	return Users
}

func (d *DatabaseLayer) GetAllProjectsFromUser(UserId string) []Models.ProjectUserDAO {
	defer d.DB.Close()

	var projectUsers []Models.ProjectUserDAO

	result, err := d.DB.Query("SELECT * FROM ProjectUser WHERE UserId = ?", UserId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var user Models.ProjectUserDAO

		err := result.Scan(&user.ProjectId, &user.UserId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		projectUsers = append(projectUsers, user)
	}

	return projectUsers
}
