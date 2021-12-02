package Test

import (
	"RokuProject-Back-End/Program"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const CreateDatabaseQuery string = "CREATE DATABASE IF NOT EXISTS RokuProjectsTest"
const CreateProjectQuery string = "CREATE TABLE IF NOT EXISTS `Project` (`Id` VARCHAR(50) NOT NULL COLLATE 'latin1_swedish_ci',`Name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`OwnerId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',PRIMARY KEY (`Id`) USING BTREE)"
const CreateProjectListQuery string = "CREATE TABLE IF NOT EXISTS `ProjectList` (`Id` VARCHAR(50) NOT NULL COLLATE 'latin1_swedish_ci',`Name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`ProjectId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`Position` INT(10,0) NULL DEFAULT NULL,PRIMARY KEY (`Id`) USING BTREE,INDEX `FK_ProjectList_Project` (`ProjectId`) USING BTREE,CONSTRAINT `FK_ProjectList_Project` FOREIGN KEY (`ProjectId`) REFERENCES `RokuProjects`.`Project` (`Id`) ON UPDATE NO ACTION ON DELETE NO ACTION)"
const CreateProjectItemQuery string = "CREATE TABLE `ProjectItem` (`Id` VARCHAR(50) NOT NULL COLLATE 'latin1_swedish_ci',`Title` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`Description` VARCHAR(255) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`ProjectListId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`Position` INT(10,0) NULL DEFAULT NULL,PRIMARY KEY (`Id`) USING BTREE,INDEX `FK_ProjectItem_ProjectList` (`ProjectListId`) USING BTREE,CONSTRAINT `FK_ProjectItem_ProjectList` FOREIGN KEY (`ProjectListId`) REFERENCES `RokuProjects`.`ProjectList` (`Id`) ON UPDATE NO ACTION ON DELETE NO ACTION)"
const CreateProjectUserQuery string = "CREATE TABLE `ProjectUser` (`ProjectId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_0900_ai_ci',`UserId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_0900_ai_ci')"
const DropDatabaseQuery string = "DROP DATABASE IF EXISTS RokuProjectsTest"

var a *Program.App

func TestMain(m *testing.M) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "locahost",
		DBName: "RokuProjectsTest",
	}

	a = Program.BuildApp(cfg)

	code := m.Run()

	DropDatabase()

	os.Exit(code)

	EnsureDatabaseCreated()

}

func EnsureDatabaseCreated() {
	if _, err := a.DB.Exec(CreateDatabaseQuery); err != nil {
		log.Fatal(err)
	}

	if _, err := a.DB.Exec(CreateProjectQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(CreateProjectListQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(CreateProjectItemQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(CreateProjectUserQuery); err != nil {
		log.Fatal(err)
	}
}

func DropDatabase() {
	if _, err := a.DB.Exec(DropDatabaseQuery); err != nil {
		log.Fatal(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
