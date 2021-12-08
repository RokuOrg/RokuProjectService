package LogicLayer

/*
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
*/
