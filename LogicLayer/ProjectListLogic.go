package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Logic"
	"github.com/segmentio/ksuid"
)

func AddProjectList(newProjectList Logic.CreateProjectList, projectId string, UserId string) (Logic.Message, string) {
	projectUser := Datalayer.GetProjectUser(Logic.ProjectUser{UserId: UserId, ProjectId: projectId})

	if projectUser.UserId == "" {
		return Logic.Message{}, "Error User not in this project"
	}

	project := Datalayer.GetProjectFromId(projectId)

	var projectList = Logic.ProjectList{
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

func GetProjectList(Id string) Logic.ProjectList {

	return Logic.ProjectList{}
}
