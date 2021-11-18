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

	Id := r.Header.Get("X-User-Validated")

	if Id == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}

	res := LogicLayer.GetProjectList(vars["id"])

	json.NewEncoder(w).Encode(res)
}

func AddProjectList(w http.ResponseWriter, r *http.Request) {

	var ProjectList Models.CreateProjectList
	vars := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&ProjectList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserId := r.Header.Get("X-User-Validated")

	if UserId == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}

	res, error := LogicLayer.AddProjectList(ProjectList, vars["project"], UserId)

	if error != "" {
		http.Error(w, error, http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func RemoveProjectList(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}

func UpdateProjectList(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
