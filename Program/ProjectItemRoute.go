package Program

import (
	"RokuProject-Back-End/Logic"
	"RokuProject-Back-End/LogicLayer"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *App) GetProjectItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Id := r.Header.Get("X-User-Validated")

	if Id == "" {
		http.Error(w, "User not validated", http.StatusUnauthorized)
		return
	}

	res := LogicLayer.GetProjectItem(vars["id"])
	json.NewEncoder(w).Encode(res)
}

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

	res, error := LogicLayer.AddProjectItem(UId, vars["project"], vars["list"], createItem)

	if error != "" {
		log.Fatal(error)
	}

	json.NewEncoder(w).Encode(res)
}

func (a *App) RemoveProjectItem(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}

func (a *App) UpdateProjectItem(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Method not implemented")
}