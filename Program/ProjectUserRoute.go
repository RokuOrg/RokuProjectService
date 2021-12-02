package Program

import (
	"RokuProject-Back-End/Logic"
	"RokuProject-Back-End/LogicLayer"
	"encoding/json"
	"net/http"
)

func (a *App) CreateProjectUser(w http.ResponseWriter, r *http.Request) {
	var ProjectUser Logic.ProjectUser

	err := json.NewDecoder(r.Body).Decode(&ProjectUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LogicLayer.CreateProjectUser(ProjectUser)

	json.NewEncoder(w).Encode(res)
}

func (a *App) RemoveUserFromProject(w http.ResponseWriter, r *http.Request) {
	var ProjectUser Logic.ProjectUser

	err := json.NewDecoder(r.Body).Decode(&ProjectUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LogicLayer.RemoveProjectUser(ProjectUser)

	json.NewEncoder(w).Encode(res)
}
