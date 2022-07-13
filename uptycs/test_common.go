package uptycs

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"net/http"
	"os"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func RespFromFixture(fixtureFile string) (*http.Response, error) {
	var fixture map[string]interface{}
	fixtureStr, err := os.ReadFile(fixtureFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fixtureStr, &fixture)
	if err != nil {
		return nil, err
	}

	resp, err := httpmock.NewJsonResponse(200, fixture)
	return resp, err
}
