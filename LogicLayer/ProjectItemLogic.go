package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Models"
	"github.com/segmentio/ksuid"
)

func AddProjectItem(UserId string, project string, list string, newItem Models.CreateProjectItem) (Models.Message, string) {
	projectUser := Datalayer.GetProjectUser(Models.ProjectUser{UserId: UserId, ProjectId: project})

	if projectUser.UserId == "" {
		return Models.Message{}, "Error User not in this project"
	}

	var projectItem = Models.ProjectItem{
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

func GetProjectItem(id string) Models.ProjectItem {
	item := Datalayer.GetProjectItem(id)

	return item
}
