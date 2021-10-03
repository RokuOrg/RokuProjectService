package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Models"
	"github.com/segmentio/ksuid"
)

func AddProjectList(newProjectList Models.CreateProjectList) Models.Message {
	project := Datalayer.GetProjectFromId(newProjectList.ProjectId)

	var projectList = Models.ProjectList{
		Id:       ksuid.New().String(),
		Name:     newProjectList.Name,
		Position: len(project.Project),
	}

	res := Datalayer.AddProjectList(projectList, newProjectList.ProjectId)

	return res
}

func RemoveProjectList() {

}

func MoveProjectItems() {

}

func GetProjectList(Id string) Models.ProjectList {

	return Models.ProjectList{}
}
