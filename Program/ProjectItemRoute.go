package Program

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (a *App) GetProjectItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Id := r.Header.Get("X-User-Validated")

	if Id == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}
	res, err := a.ProjectCollection.GetProject(vars["id"])

	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
	}

	e := json.NewEncoder(w).Encode(res)
	if e != nil {
		return
	}
}

/*
func (a *App) AddProjectItem(w http.ResponseWriter, r *http.Request) {
	var createItem Logic.CreateProjectItem

	err := json.NewDecoder(r.Body).Decode(&createItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)

	UId := r.Header.Get("X-User-Validated")

	if UId == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
	}

	res, e := LogicLayer.AddProjectItem(UId, vars["project"], vars["list"], createItem)

	if e != "" {
		log.Fatal(e)
	}

	json.NewEncoder(w).Encode(res)
}

func (a *App) RemoveProjectItem(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
*/
/*
func (a *App) UpdateProjectItem(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}
*/
