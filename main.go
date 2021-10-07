package main

import (
	"RokuProject-Back-End/Routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {

	r := mux.NewRouter().StrictSlash(true)

	//Project
	r.HandleFunc("/project/all", Routes.ProjectsPage).Methods("Get") // done

	r.HandleFunc("/project", Routes.CreateProject).Methods("Post") //test

	r.HandleFunc("/project", Routes.RemoveProject).Methods("Delete") // todo

	r.HandleFunc("/project", Routes.UpdateProject).Methods("Update") // todo

	r.HandleFunc("/project/{id}", Routes.GetProjectById).Methods("Get") // test

	r.HandleFunc("/project/user", Routes.GetProjectsByUserId).Methods("Get") // test

	//Project User
	r.HandleFunc("/{project}/user", Routes.RemoveUserFromProject).Methods("Delete") //todo

	r.HandleFunc("/project/user", Routes.CreateProjectUser).Methods("Post") // todo

	//Project Item
	r.HandleFunc("/{project}/{list}/item", Routes.AddProjectItem).Methods("Post")

	r.HandleFunc("/{project}/{list}/{item}", Routes.GetProjectItem).Methods("Get")

	r.HandleFunc("/{project}/{list}/{item}", Routes.RemoveProjectItem).Methods("Delete")

	r.HandleFunc("/{project}/{list}/{item}", Routes.UpdateProjectItem).Methods("Update")

	//Project List
	r.HandleFunc("/{project}/list", Routes.AddProjectList).Methods("Post")

	r.HandleFunc("/{project}/{list}", Routes.GetProjectList).Methods("Get")

	r.HandleFunc("/{project}/{list}", Routes.RemoveProjectList).Methods("Delete")

	r.HandleFunc("/{project}/{list}", Routes.UpdateProjectList).Methods("Update")

	//Middlewares

	port := ":7001"

	fmt.Println("Listening on Port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	HandleRequest()
}
