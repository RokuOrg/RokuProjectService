package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Logic"
	"log"
)

func CreateProjectUser(user Logic.ProjectUser) Logic.Message {
	res := Datalayer.AddProjectUser(user)

	return res
}

func RemoveProjectUser(user Logic.ProjectUser) Logic.Message {
	res := Datalayer.RemoveProjectUser(user)

	return res
}

func GetProjectUser(ProjectUser Logic.ProjectUser) Logic.ProjectUser {
	user := Datalayer.GetProjectUser(ProjectUser)

	return user
}

//Functions under this comment are for Multiple ProjectUsers

func RemoveProjectUsersFromProjectId() Logic.Message {
	log.Fatal("method not implemented")
	return Logic.Message{}
}

func RemoveProjectUsersFromUserId() Logic.Message {
	log.Fatal("method not implemented")
	return Logic.Message{}
}

func GetProjectUsersFromProjectId() Logic.ProjectUsers {
	log.Fatal("method not implemented")
	return Logic.ProjectUsers{}
}

func GetProjectUsersFromUserId() Logic.ProjectUsers {
	log.Fatal("method not implemented")
	return Logic.ProjectUsers{}
}
