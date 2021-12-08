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

	rows := sqlmock.NewRows([]string{"Id", "Name", "OwnerId"})

	mock.ExpectQuery("INSERT INTO Project").WithArgs(FakeRandomId().String(), "test project", "12345").WillReturnRows(rows)
	//mock Add project list
	//mock Add project items
	//mock Add Project user
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
