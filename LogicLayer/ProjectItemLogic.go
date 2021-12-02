package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Logic"
	"github.com/segmentio/ksuid"
)

func AddProjectItem(UserId string, project string, list string, newItem Logic.CreateProjectItem) (Logic.Message, string) {
	projectUser := Datalayer.GetProjectUser(Logic.ProjectUser{UserId: UserId, ProjectId: project})

	if projectUser.UserId == "" {
		return Logic.Message{}, "Error User not in this project"
	}

	var projectItem = Logic.ProjectItem{
		Position:    newItem.Position,
		Title:       newItem.Title,
		Description: "",
		Id:          ksuid.New().String(),
	}

	res := Datalayer.AddProjectItem(projectItem, list)

	return res, ""
}

func RemoveProjectItem() {

}

func GetProjectItem(id string) Logic.ProjectItem {
	item := Datalayer.GetProjectItem(id)

	return item
}
