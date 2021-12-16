package Test

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"testing"
)

func TestCreateProject(t *testing.T) {

	//Setup
	var jsonStr = []byte(`{"name": "test project", "template": true}`)

	req, _ := http.NewRequest("POST", "/project", bytes.NewBuffer(jsonStr))
	req.Header.Add("X-User-Validated", "12345")

	//mock Add project
	//, project.Id, project.Name, project.OwnerId

	ProjectRows := sqlmock.NewRows([]string{"Id", "Name", "OwnerId"})
	ProjectListRows := sqlmock.NewRows([]string{"Id", "Name", "ProjectId", "Position"})
	//ProjectItemRows := sqlmock.NewRows([]string{"Id", "Title", "Description", "ProjectListId", "Position"})
	ProjectUserRows := sqlmock.NewRows([]string{"ProjectId", "UserId"})

	mock.ExpectQuery("INSERT INTO Project").WithArgs(FakeRandomId().String(), "test project", "12345").WillReturnRows(ProjectRows)

	//mock Add project list
	mock.ExpectQuery("INSERT INTO ProjectList").WithArgs(FakeRandomId().String(), "To Do", FakeRandomId().String(), 0).WillReturnRows(ProjectListRows)
	mock.ExpectQuery("INSERT INTO ProjectList").WithArgs(FakeRandomId().String(), "In Progress", FakeRandomId().String(), 1).WillReturnRows(ProjectListRows)
	mock.ExpectQuery("INSERT INTO ProjectList").WithArgs(FakeRandomId().String(), "Done", FakeRandomId().String(), 2).WillReturnRows(ProjectListRows)

	//mock Add Project user
	mock.ExpectQuery("INSERT INTO ProjectUser").WithArgs(FakeRandomId().String(), "12345").WillReturnRows(ProjectUserRows)
	defer mock.ExpectClose()

	//Execute
	res := executeRequest(req)

	//Validate
	if res.Code != http.StatusOK {
		t.Errorf("StatusCode was: %v, Response: %s", res.Code, res.Body)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s", err)
	}

}
