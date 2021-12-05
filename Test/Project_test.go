package Test

import (
	"bytes"
	"net/http"
	"testing"
)

func TestCreateProject(t *testing.T) {

	var jsonStr = []byte(`{"name": "test project", "template": true}`)

	req, _ := http.NewRequest("POST", "/project", bytes.NewBuffer(jsonStr))
	req.Header.Add("X-User-Validated", "12345")
	res := executeRequest(req)

	if res.Code != http.StatusOK {
		t.Errorf("StatusCode was: %v, Response: %s", res.Code, res.Body)
	}

	result, err := a.DB.Query("SELECT * FROM Project WHERE OwnerId = 12345")

	if err != nil {
		t.Errorf("SQL Error: %s", err)
	}

	for result.Next() {

	}

}
