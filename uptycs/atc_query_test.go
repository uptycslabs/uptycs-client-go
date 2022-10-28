package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAtcQuery(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		id      string
		out     interface{}
	}

	theTests := []convTest{
		{
			name:    "TestAtcQuery",
			fixture: "fixtures/atcQuery.json",
			id:      "dc0e9652-ec9a-4baa-9da3-8547333b3628",
			out: AtcQuery{
				Name:        "atc_chrome_browser_history",
				Description: "Chrome Browser History",
				Query:       "SELECT urls.id id, urls.url url, urls.title title, urls.visit_count visit_count, urls.typed_count typed_count, urls.last_visit_time last_visit_time, urls.hidden hidden, visits.visit_time visit_time, visits.from_visit from_visit, visits.visit_duration visit_duration, visits.transition transition, visit_source.source source FROM urls JOIN visits ON urls.id = visits.url LEFT JOIN visit_source ON visits.id = visit_source.id",
				OsPaths: struct {
					Darwin  []PathStruct `json:"darwin,omitempty"`
					Debian  []PathStruct `json:"debian,omitempty"`
					Windows []PathStruct `json:"windows,omitempty"`
				}{
					Darwin:  []PathStruct{{Path: "/Users/%/Library/Application Support/Google/Chrome/%/History"}},
					Debian:  []PathStruct{{Path: "/home/%/.config/google-chrome/Default/History"}},
					Windows: []PathStruct{{Path: "C:\\Users\\%\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\History"}},
				},
				Columns: []struct {
					Name        string `json:"name,omitempty"`
					Description string `json:"description,omitempty"`
				}{
					{Name: "id", Description: "Id"},
					{Name: "path", Description: "Path"},
					{Name: "url", Description: "Url"},
					{Name: "visit_count", Description: "VisitCount"},
					{Name: "typed_count", Description: "TypedCount"},
					{Name: "last_visit_time", Description: "LastVisitTime"},
					{Name: "hidden", Description: "Hidden"},
					{Name: "visit_time", Description: "VisitTime"},
					{Name: "visit_duration", Description: "VisitDuration"},
					{Name: "source", Description: "Source"},
				},
				CreatedBy: "00000000-0000-0000-0000-000000000000",
				UpdatedBy: "00000000-0000-0000-0000-000000000000",
				CreatedAt: "2022-01-17T08:12:16.674Z",
				UpdatedAt: "2022-10-27T03:51:08.974Z", ID: "dc0e9652-ec9a-4baa-9da3-8547333b3628",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "ATC query information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/atcQueries/dc0e9652-ec9a-4baa-9da3-8547333b3628"},
					LinkItem{Rel: "parent", Title: "ATC queries information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/atcQueries"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/atcQueries/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			atcQueryResp, err := c.GetAtcQuery(AtcQuery{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(atcQueryResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", atcQueryResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteAtcQuery(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   AtcQuery
	}

	theTests := []convTest{
		{
			name: "TestAtcQuery",
			in: AtcQuery{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/atcQueries/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteAtcQuery(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/atcQueries/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutAtcQuery(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AtcQuery
	}

	theTests := []convTest{
		{
			name:    "TestAtcQuery",
			fixture: "fixtures/atcQuery.json",
			in: AtcQuery{
				ID: "dc0e9652-ec9a-4baa-9da3-8547333b3628",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "ATC query information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/atcQueries/dc0e9652-ec9a-4baa-9da3-8547333b3628"},
					LinkItem{Rel: "parent", Title: "ATC queries information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/atcQueries"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/atcQueries/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateAtcQuery(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/atcQueries/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateAtcQuery(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AtcQuery
	}

	theTests := []convTest{
		{
			name:    "TestAtcQuery",
			fixture: "fixtures/atcQuery.json",
			in: AtcQuery{
				ID: "dc0e9652-ec9a-4baa-9da3-8547333b3628",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/atcQueries",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateAtcQuery(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/atcQueries"], 1)
		})
	}
}
