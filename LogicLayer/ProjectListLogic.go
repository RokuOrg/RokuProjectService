package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Models"
	"github.com/segmentio/ksuid"
)

func AddProjectList(newProjectList Models.CreateProjectList, projectId string, UserId string) (Models.Message, string) {
	projectUser := Datalayer.GetProjectUser(Models.ProjectUser{UserId: UserId, ProjectId: projectId})

	if projectUser.UserId == "" {
		return Models.Message{}, "Error User not in this project"
	}

	project := Datalayer.GetProjectFromId(projectId)

	var projectList = Models.ProjectList{
		Id:       ksuid.New().String(),
		Name:     newProjectList.Name,
		Position: len(project.Project),
	}

	res := Datalayer.AddProjectList(projectList, projectId)

	return res, ""
}

func RemoveProjectList() {

}

func MoveProjectItems() {

}

func GetProjectList(Id string) Models.ProjectList {

	return Models.ProjectList{}
}
