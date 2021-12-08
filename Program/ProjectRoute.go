package Program

import (
	"RokuProject-Back-End/Models"
	"encoding/json"
	"log"
	"net/http"
)

/*
func (a *App) ProjectsPage(w http.ResponseWriter, r *http.Request) {
	projects := LogicLayer.GetAllProjects()

	json.NewEncoder(w).Encode(projects)
}
*/
func (a *App) CreateProject(w http.ResponseWriter, r *http.Request) {
	var createProject Models.ProjectDTO

	uId := r.Header.Get("X-User-Validated")

	if uId == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
	}

	err := json.NewDecoder(r.Body).Decode(&createProject)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.ProjectCollection.CreateProject(uId, createProject.Name, createProject.Template); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(http.StatusOK)
}

/*
func (a *App) GetProjectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Id := r.Header.Get("X-User-Validated")

	if Id == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}

	project, err := LogicLayer.GetProjectByUser(vars["id"], Id)

	if err != "" {
		http.Error(w, err, http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func (a *App) GetProjectsByUserId(w http.ResponseWriter, r *http.Request) {
	Id := r.Header.Get("X-User-Validated")

	if Id == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}

	projects := LogicLayer.GetProjectsFromUser(Id)

	json.NewEncoder(w).Encode(projects)
}
*/
func (a *App) RemoveProject(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}

func (a *App) UpdateProject(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
