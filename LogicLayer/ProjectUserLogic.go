package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Models"
	"log"
)

func CreateProjectUser(user Models.ProjectUser) Models.Message {
	res := Datalayer.AddProjectUser(user)

	return res
}

func RemoveProjectUser(user Models.ProjectUser) Models.Message {
	res := Datalayer.RemoveProjectUser(user)

	return res
}

func GetProjectUser(ProjectUser Models.ProjectUser) Models.ProjectUser {
	user := Datalayer.GetProjectUser(ProjectUser)

	return user
}

//Functions under this comment are for Multiple ProjectUsers

func RemoveProjectUsersFromProjectId() Models.Message {
	log.Fatal("method not implemented")
	return Models.Message{}
}

func RemoveProjectUsersFromUserId() Models.Message {
	log.Fatal("method not implemented")
	return Models.Message{}
}

func GetProjectUsersFromProjectId() Models.ProjectUsers {
	log.Fatal("method not implemented")
	return Models.ProjectUsers{}
}

func GetProjectUsersFromUserId() Models.ProjectUsers {
	log.Fatal("method not implemented")
	return Models.ProjectUsers{}
}
