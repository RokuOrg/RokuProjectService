package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Models"
	"github.com/segmentio/ksuid"
)

func ProjectListTemplate() Models.ProjectLists {

	ProjectList := Models.ProjectLists{
		Models.ProjectList{
			Id:       ksuid.New().String(),
			Name:     "To Do",
			Position: 0,
			Items:    Models.ProjectItems{},
		},
		Models.ProjectList{
			Id:       ksuid.New().String(),
			Name:     "In Progress",
			Position: 1,
			Items:    Models.ProjectItems{},
		},
		Models.ProjectList{
			Id:       ksuid.New().String(),
			Name:     "Done",
			Position: 2,
			Items:    Models.ProjectItems{},
		},
	}

	return ProjectList
}

func CreateProject(UserId string, ProjectName string, template bool) Models.Message {

	project := Models.Project{
		Name:    ProjectName,
		Id:      ksuid.New().String(),
		OwnerId: UserId,
		Project: nil,
	}

	if template {
		project.Project = ProjectListTemplate()
	}

	res := Datalayer.AddProject(project)

	if !res.Success {
		return res
	}

	for i := range project.Project {
		res := Datalayer.AddProjectList(project.Project[i], project.Id)

		if !res.Success {
			return res
		}
		for j := range project.Project[i].Items {
			res := Datalayer.AddProjectItem(project.Project[i].Items[j], project.Project[i].Id)

			if !res.Success {
				return res
			}
		}
	}

	res = Datalayer.AddProjectUser(Models.ProjectUser{UserId: UserId, ProjectId: project.Id})

	return res
}

func GetProjectByUser(ProjectId string, UserId string) (Models.Project, string) {

	ProjectUser := Datalayer.GetProjectUser(Models.ProjectUser{UserId: UserId, ProjectId: ProjectId})

	if ProjectUser.ProjectId == "" {
		return Models.Project{}, "User not part of project"
	}

	project := Datalayer.GetProjectFromId(ProjectId)

	project.Project = Datalayer.GetProjectListsFromProjectId(ProjectId)

	for i := range project.Project {
		project.Project[i].Items = Datalayer.GetProjectItemsFromProjectList(project.Project[i].Id)
	}

	return project, ""
}

func GetProject(ProjectId string) Models.Project {

	project := Datalayer.GetProjectFromId(ProjectId)

	project.Project = Datalayer.GetProjectListsFromProjectId(ProjectId)

	for i := range project.Project {
		project.Project[i].Items = Datalayer.GetProjectItemsFromProjectList(project.Project[i].Id)
	}

	return project
}

func MoveProjectLists() {

}

//Functions under this comment are for Multiple Projects

func GetProjectsFromUser(userId string) Models.Projects {
	projectUsers := Datalayer.GetAllProjectsFromUser(userId)
	var projects Models.Projects

	for i := range projectUsers {
		projects = append(projects, GetProject(projectUsers[i].ProjectId))
	}
	return projects
}

func GetAllProjects() Models.Projects {
	project1 := Datalayer.GetAllProjects()
	var projects Models.Projects

	for i := range project1 {
		projects = append(projects, GetProject(project1[i].Id))
	}

	return projects
}
