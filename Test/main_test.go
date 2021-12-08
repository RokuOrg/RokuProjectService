package Test

import (
	"RokuProject-Back-End/Program"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/segmentio/ksuid"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const CreateProjectQuery string = "CREATE TABLE IF NOT EXISTS `Project` (`Id` VARCHAR(50) NOT NULL COLLATE 'latin1_swedish_ci',`Name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',`OwnerId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',PRIMARY KEY (`Id`) USING BTREE) COLLATE='latin1_swedish_ci'ENGINE=InnoDB;"
const CreateProjectListQuery string = "CREATE TABLE IF NOT EXISTS `ProjectList` (\n\t`Id` VARCHAR(50) NOT NULL COLLATE 'latin1_swedish_ci',\n\t`Name` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`ProjectId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`Position` INT(10) NULL DEFAULT NULL,\n\tPRIMARY KEY (`Id`) USING BTREE,\n\tINDEX `FK_ProjectList_Project` (`ProjectId`) USING BTREE,\n\tCONSTRAINT `FK_ProjectList_Project` FOREIGN KEY (`ProjectId`) REFERENCES `RokuProjects`.`Project` (`Id`) ON UPDATE NO ACTION ON DELETE NO ACTION\n)\nCOLLATE='latin1_swedish_ci'\nENGINE=InnoDB\n;"
const CreateProjectItemQuery string = "CREATE TABLE IF NOT EXISTS `ProjectItem` (\n\t`Id` VARCHAR(50) NOT NULL COLLATE 'latin1_swedish_ci',\n\t`Title` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`Description` VARCHAR(255) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`ProjectListId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',\n\t`Position` INT(10) NULL DEFAULT NULL,\n\tPRIMARY KEY (`Id`) USING BTREE,\n\tINDEX `FK_ProjectItem_ProjectList` (`ProjectListId`) USING BTREE,\n\tCONSTRAINT `FK_ProjectItem_ProjectList` FOREIGN KEY (`ProjectListId`) REFERENCES `RokuProjects`.`ProjectList` (`Id`) ON UPDATE NO ACTION ON DELETE NO ACTION\n)\nCOLLATE='latin1_swedish_ci'\nENGINE=InnoDB\n;"
const CreateProjectUserQuery string = "CREATE TABLE IF NOT EXISTS `ProjectUser` (\n\t`ProjectId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_0900_ai_ci',\n\t`UserId` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_0900_ai_ci'\n)\nCOLLATE='utf8mb4_0900_ai_ci'\nENGINE=InnoDB\n;"

var a *Program.App
var mock sqlmock.Sqlmock

func FakeRandomId() ksuid.KSUID {
	id, err := ksuid.Parse("21v9w2MFx7wOoWt7wjfTohx5viA")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return id
}

func TestMain(m *testing.M) {

	db, M, err := sqlmock.New()
	mock = M
	if err != nil {
		fmt.Println("expected no error, but got:", err)
		return
	}

	defer db.Close()

	a = Program.BuildApp(db, FakeRandomId)

	//EnsureDatabaseCreated()

	code := m.Run()

	os.Exit(code)
}

func EnsureDatabaseCreated() {
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
