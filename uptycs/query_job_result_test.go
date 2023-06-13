package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetQueryJobResult(t *testing.T) {

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
			name:    "TestQueryJobResult",
			fixture: "fixtures/query_job_result.json",
			id:      "7bc5024f-003b-4934-97d6-478bc603a684",
			out: QueryJobResult{
				QueryStats: struct {
					CPUTimeMillis     int `json:"cpuTimeMillis"`
					ProcessedRows     int `json:"processedRows"`
					ProcessedBytes    int `json:"processedBytes"`
					ElapsedTimeMillis int `json:"elapsedTimeMillis"`
				}{
					CPUTimeMillis:     106,
					ProcessedRows:     5963,
					ProcessedBytes:    0,
					ElapsedTimeMillis: 386,
				},
				Status:      "FINISHED",
				ID:          "7bc5024f-003b-4934-97d6-478bc603a684",
				Name:        "",
				RowDataHash: "",
				Error:       nil,
				EndTime:     "2023-06-13T15:27:01.209Z",
				StartTime:   "2023-06-13T15:27:00.692Z",
				RowCount:    2,
				ResultStore: "cache",
				RowData:     "",
				RowNumber:   0,
				QueryJobID:  "",
				Columns: []QueryJobColumn{
					{
						Name:         "instance_id",
						Type:         "TEXT",
						OriginalName: "instance_id",
					},
					{
						Name:         "tags",
						Type:         "TEXT",
						OriginalName: "tags",
					},
					{
						Name:         "foo",
						Type:         "NUMBER_INTEGER",
						OriginalName: "foo",
					},
				},
				Offset: 0,
				Limit:  50000,
				Items: []QueryJobResult{
					{
						QueryStats: struct {
							CPUTimeMillis     int `json:"cpuTimeMillis"`
							ProcessedRows     int `json:"processedRows"`
							ProcessedBytes    int `json:"processedBytes"`
							ElapsedTimeMillis int `json:"elapsedTimeMillis"`
						}{
							CPUTimeMillis:     0,
							ProcessedRows:     0,
							ProcessedBytes:    0,
							ElapsedTimeMillis: 0,
						},
						Status:      "",
						ID:          "",
						Name:        "",
						RowDataHash: "248f094d-c546-34dc-8ca2-06c869c4bbed",
						Error:       nil,
						CreatedAt:   "2023-06-13",
						EndTime:     "",
						StartTime:   "",
						RowCount:    0,
						RowData:     "{\"foo\":1,\"instance_id\":\"i-027e4ca22fe5a1518\",\"tags\":\"{\\\"Foo\\\":\\\"some foo\\\"}\"}",
						ResultStore: "",
						RowNumber:   1,
						QueryJobID:  "7bc5024f-003b-4934-97d6-478bc603a684",
						Columns:     nil,
						Offset:      0,
						Limit:       0,
						Items:       nil,
						Links:       nil,
					},
					{
						QueryStats: struct {
							CPUTimeMillis     int `json:"cpuTimeMillis"`
							ProcessedRows     int `json:"processedRows"`
							ProcessedBytes    int `json:"processedBytes"`
							ElapsedTimeMillis int `json:"elapsedTimeMillis"`
						}{
							CPUTimeMillis:     0,
							ProcessedRows:     0,
							ProcessedBytes:    0,
							ElapsedTimeMillis: 0,
						},
						Status:      "",
						ID:          "",
						Name:        "",
						RowDataHash: "8917663c-4f37-3d28-b771-acfd00f2e816",
						Error:       nil,
						CreatedAt:   "2023-06-13",
						EndTime:     "",
						StartTime:   "",
						RowCount:    0,
						RowData:     "{\"foo\":2,\"instance_id\":\"i-08c0f696e648e37cd\",\"tags\":\"{\\\"Foo\\\":\\\"some other foo\\\"}\"}",
						ResultStore: "",
						RowNumber:   2,
						QueryJobID:  "7bc5024f-003b-4934-97d6-478bc603a684",
						Columns:     nil,
						Offset:      0,
						Limit:       0,
						Items:       nil,
						Links:       nil,
					},
				},
				Links: []LinkItem{
					{
						Rel:   "self",
						Title: "Query job results information",
						Href:  "/api/customers/11111111-1111-1111-1111-111111111111/queryJobs/7bc5024f-003b-4934-97d6-478bc603a684/results",
					},
					{
						Rel:   "parent",
						Title: "Query job information",
						Href:  "/api/customers/11111111-1111-1111-1111-111111111111/queryJobs/7bc5024f-003b-4934-97d6-478bc603a684",
					},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/queryJobs/%v/results", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			queryJobResp, err := c.GetQueryJobResults(QueryJobResult{
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

func TestDeleteQueryJobResult(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   QueryJobResult
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

			_, err := c.DeleteQueryJobResult(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/queryJobs/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutQueryJobResult(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      QueryJobResult
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

			_, err := c.UpdateQueryJobResults(theT.in)

			expectedErrorMsg := "UPDATE is not supported for lookup tables"
			assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
		})
	}
}

func TestCreateQueryJobResult(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      QueryJobResult
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

			_, err := c.CreateQueryJobResults(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/queryJobs"], 1)
		})
	}
}
