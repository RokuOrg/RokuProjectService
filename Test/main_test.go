package Test

import (
	"RokuProject-Back-End/Program"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/segmentio/ksuid"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
