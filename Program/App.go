package Program

import (
	"RokuProject-Back-End/Logic"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *App) Run(port string) {
	fmt.Println("Listening on Port " + port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

type App struct {
	Router            *mux.Router
	DB                *sql.DB
	ProjectCollection *Logic.ProjectCollection
}

func BuildApp(cfg mysql.Config) *App {
	a := App{}

	var query string = cfg.User + ":" + cfg.Passwd + "@" + cfg.Net + "(" + cfg.Addr + ")/"

	db, er := sql.Open("mysql", query)

	if er != nil {
		log.Fatal(er)
	}

	_, er = db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.DBName)

	if er != nil {
		log.Fatal(er)
	}

	var err error
	a.DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := a.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	a.Router = mux.NewRouter().StrictSlash(true)

	//Project
	a.Router.HandleFunc("/project/all", a.ProjectsPage).Methods("Get") // done

	a.Router.HandleFunc("/project", a.CreateProject).Methods("Post") //test

	a.Router.HandleFunc("/project", a.RemoveProject).Methods("Delete") // todo

	a.Router.HandleFunc("/project", a.UpdateProject).Methods("Update") // todo

	//Project User
	a.Router.HandleFunc("/project/user", a.GetProjectsByUserId).Methods("Get") // test

	a.Router.HandleFunc("/project/user", a.RemoveUserFromProject).Methods("Delete") //todo

	a.Router.HandleFunc("/project/user", a.CreateProjectUser).Methods("Post") // todo

	//Project List
	a.Router.HandleFunc("/{project}/list", a.AddProjectList).Methods("Post") //test

	a.Router.HandleFunc("/{project}/{list}", a.GetProjectList).Methods("Get") //test

	a.Router.HandleFunc("/{project}/{list}", a.RemoveProjectList).Methods("Delete") //todo

	a.Router.HandleFunc("/{project}/{list}", a.UpdateProjectList).Methods("Update") //todo

	//Project Item
	a.Router.HandleFunc("/{project}/{list}/item", a.AddProjectItem).Methods("Post") //test

	a.Router.HandleFunc("/{project}/{list}/{item}", a.GetProjectItem).Methods("Get") //test

	a.Router.HandleFunc("/{project}/{list}/{item}", a.RemoveProjectItem).Methods("Delete") //todo

	a.Router.HandleFunc("/{project}/{list}/{item}", a.UpdateProjectItem).Methods("Update") //todo

	a.Router.HandleFunc("/project/{id}", a.GetProjectById).Methods("Get") // test

	a.ProjectCollection = Logic.BuildProjectCollection(a.DB)

	return &a
}
