package Test

import (
	"net/http"
	"testing"
)

func TestCreateProject(t *testing.T) {
	req, _ := http.NewRequest("POST", "/project", nil)
	res := executeRequest(req)

	if res.Code != http.StatusOK {
		t.Errorf("StatusCode was not OK : %v", res.Code)
	}

}
