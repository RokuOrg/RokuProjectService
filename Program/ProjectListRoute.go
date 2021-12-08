package Program

import (
	"log"
	"net/http"
)

/*
func (a *App) GetProjectList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Id := r.Header.Get("X-User-Validated")

	if Id == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}

	res := LogicLayer.GetProjectList(vars["id"])

	json.NewEncoder(w).Encode(res)
}

func (a *App) AddProjectList(w http.ResponseWriter, r *http.Request) {

	var ProjectList Logic.CreateProjectList
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
*/
func (a *App) RemoveProjectList(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}

func (a *App) UpdateProjectList(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
