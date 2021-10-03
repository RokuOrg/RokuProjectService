package Routes

import (
	"RokuProject-Back-End/LogicLayer"
	"RokuProject-Back-End/Models"
	"encoding/json"
	"net/http"
)

func CreateProjectUser(w http.ResponseWriter, r *http.Request) {
	var ProjectUser Models.ProjectUser

	err := json.NewDecoder(r.Body).Decode(&ProjectUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LogicLayer.CreateProjectUser(ProjectUser)

	json.NewEncoder(w).Encode(res)
}

func RemoveUserFromProject(w http.ResponseWriter, r *http.Request) {
	var ProjectUser Models.ProjectUser

	err := json.NewDecoder(r.Body).Decode(&ProjectUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LogicLayer.RemoveProjectUser(ProjectUser)

	json.NewEncoder(w).Encode(res)
}
