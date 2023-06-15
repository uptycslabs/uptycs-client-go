package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetQueryJob(t *testing.T) {

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
			name:    "TestQueryJob",
			fixture: "fixtures/query_job.json",
			id:      "b6885999-3bb0-4d8e-84b3-46883f4d7506",
			out: QueryJob{
				ID:    "b6885999-3bb0-4d8e-84b3-46883f4d7506",
				Name:  "b86c2befb58987cb5d4f14210784c9df",
				Query: "test",
				Type:  "global",
				Parameters: []QueryJobParameter{
					{
						Key:          "from",
						DataType:     "DATE",
						Multiple:     false,
						Optional:     false,
						DefaultValue: "2023-05-22T17:22:36.323Z",
					},
					{
						Key:          "to",
						DataType:     "DATE",
						Multiple:     false,
						Optional:     false,
						DefaultValue: "2023-05-22T19:27:08.322Z",
					},
				},
				ParameterValues: struct {
					From string `json:"from,omitempty"`
					To   string `json:"to,omitempty"`
				}{
					From: "2023-05-22T17:22:36.323Z",
					To:   "2023-05-22T19:27:08.322Z",
				},
				QueryID:           "4934a4b8-0625-452f-b5ec-0d6be1da2d85",
				Status:            "QUEUED",
				RowCount:          0,
				Columns:           nil,
				StartTime:         "",
				EndTime:           "",
				Error:             QueryError{},
				Purged:            false,
				IncompleteResults: false,
				AlertID:           "",
				CreatedBy:         "f48f4c40-9c4a-47bb-9e3f-797d4deca92a",
				UpdatedBy:         "",
				CreatedAt:         "2023-05-22T19:27:08.334Z",
				UpdatedAt:         "2023-05-22T19:27:08.334Z",
				Source:            "SCHEDULED",
				ResultStore:       "",
				AgentType:         "asset",
				ResourceType:      "asset",
				Links: []LinkItem{
					{
						Rel:   "self",
						Title: "Query job information",
						Href:  "/api/customers/11111111-1111-1111-1111-111111111111/queryJobs/b6885999-3bb0-4d8e-84b3-46883f4d7506",
					},
					{
						Rel:   "parent",
						Title: "Query jobs information",
						Href:  "/api/customers/11111111-1111-1111-1111-111111111111/queryJobs",
					},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/queryJobs/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			queryJobResp, err := c.GetQueryJob(QueryJob{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(queryJobResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", queryJobResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteQueryJob(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   QueryJob
	}

	theTests := []convTest{}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/queryJobs/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteQueryJob(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/queryJobs/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutQueryJob(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      QueryJob
	}

	theTests := []convTest{}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/queryJobs/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateQueryJob(theT.in)

			expectedErrorMsg := "UPDATE is not supported for lookup tables"
			assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
		})
	}
}

func TestCreateQueryJob(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      QueryJob
	}

	theTests := []convTest{}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/queryJobs",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateQueryJob(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/queryJobs"], 1)
		})
	}
}
