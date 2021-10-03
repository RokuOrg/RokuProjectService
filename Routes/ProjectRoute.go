package Routes

import (
	"RokuProject-Back-End/LogicLayer"
	"RokuProject-Back-End/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ProjectsPage(w http.ResponseWriter, r *http.Request) {
	projects := LogicLayer.GetAllProjects()

	json.NewEncoder(w).Encode(projects)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var createProject Models.CreateProject

	err := json.NewDecoder(r.Body).Decode(&createProject)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LogicLayer.CreateProject(createProject.Token, createProject.Name, createProject.Template)

	json.NewEncoder(w).Encode(res)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	project := LogicLayer.GetProject(vars["id"])

	json.NewEncoder(w).Encode(project)
}

func GetProjectsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projects := LogicLayer.GetProjectsFromUser(vars["id"])

	json.NewEncoder(w).Encode(projects)
}

func RemoveProject(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
