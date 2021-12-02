package LogicLayer

import (
	"RokuProject-Back-End/Datalayer"
	"RokuProject-Back-End/Logic"
	"github.com/segmentio/ksuid"
)

func ProjectListTemplate() Logic.ProjectLists {

	ProjectList := Logic.ProjectLists{
		Logic.ProjectList{
			Id:       ksuid.New().String(),
			Name:     "To Do",
			Position: 0,
			Items:    Logic.ProjectItems{},
		},
		Logic.ProjectList{
			Id:       ksuid.New().String(),
			Name:     "In Progress",
			Position: 1,
			Items:    Logic.ProjectItems{},
		},
		Logic.ProjectList{
			Id:       ksuid.New().String(),
			Name:     "Done",
			Position: 2,
			Items:    Logic.ProjectItems{},
		},
	}

	return ProjectList
}

func CreateProject(UserId string, ProjectName string, template bool) Logic.Message {

	project := Logic.Project{
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

	res = Datalayer.AddProjectUser(Logic.ProjectUser{UserId: UserId, ProjectId: project.Id})

	return res
}

func GetProjectByUser(ProjectId string, UserId string) (Logic.Project, string) {

	ProjectUser := Datalayer.GetProjectUser(Logic.ProjectUser{UserId: UserId, ProjectId: ProjectId})

	if ProjectUser.ProjectId == "" {
		return Logic.Project{}, "User not part of project"
	}

	project := Datalayer.GetProjectFromId(ProjectId)

	project.Project = Datalayer.GetProjectListsFromProjectId(ProjectId)

	for i := range project.Project {
		project.Project[i].Items = Datalayer.GetProjectItemsFromProjectList(project.Project[i].Id)
	}

	return project, ""
}

func GetProject(ProjectId string) Logic.Project {

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

func GetProjectsFromUser(userId string) Logic.Projects {
	projectUsers := Datalayer.GetAllProjectsFromUser(userId)
	var projects Logic.Projects

	for i := range projectUsers {
		projects = append(projects, GetProject(projectUsers[i].ProjectId))
	}
	return projects
}

func GetAllProjects() Logic.Projects {
	project1 := Datalayer.GetAllProjects()
	var projects Logic.Projects

	for i := range project1 {
		projects = append(projects, GetProject(project1[i].Id))
	}

	return projects
}
