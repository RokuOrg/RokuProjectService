package Routes

import (
	"RokuProject-Back-End/LogicLayer"
	"RokuProject-Back-End/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetProjectList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res := LogicLayer.GetProjectList(vars["id"])

	json.NewEncoder(w).Encode(res)
}

func AddProjectList(w http.ResponseWriter, r *http.Request) {
	var ProjectList Models.CreateProjectList

	err := json.NewDecoder(r.Body).Decode(&ProjectList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LogicLayer.AddProjectList(ProjectList)

	json.NewEncoder(w).Encode(res)
}

func RemoveProjectList(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}

func UpdateProjectList(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
