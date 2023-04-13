package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLookupTable(t *testing.T) {

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
			name:    "TestLookupTable",
			fixture: "fixtures/lookup_table.json",
			id:      "7f042e5a-e4e8-47e7-8077-a0581411d4fd",
			out: LookupTable{
				ID:              "7f042e5a-e4e8-47e7-8077-a0581411d4fd",
				Name:            "aws_account_metadata",
				Description:     "Used to lookup business context on AWS accounts.",
				Active:          true,
				IDField:         "id",
				RowCount:        3,
				ForRuleEngine:   true,
				CreatedBy:       "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy:       "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				CreatedAt:       "2022-06-22T03:58:40.114Z",
				UpdatedAt:       "2022-06-22T03:58:40.349Z",
				DataLookupTable: DataLookupTable{},
				FetchRowsquery:  "SELECT id_field_value,data FROM upt_lookup_rows WHERE lookup_table_id = '7f042e5a-e4e8-47e7-8077-a0581411d4fd'",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Lookup Table information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/lookupTables/7f042e5a-e4e8-47e7-8077-a0581411d4fd"},
					LinkItem{Rel: "parent", Title: "Lookup Tables information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/lookupTables"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/lookupTables/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			lookupTableResp, err := c.GetLookupTable(LookupTable{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(lookupTableResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", lookupTableResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteLookupTable(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   LookupTable
	}

	theTests := []convTest{
		{
			name: "TestLookupTable",
			in: LookupTable{
				ID: "7f042e5a-e4e8-47e7-8077-a0581411d4fd",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/lookupTables/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteLookupTable(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/lookupTables/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutLookupTable(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      LookupTable
	}

	theTests := []convTest{
		{
			name:    "TestLookupTable",
			fixture: "fixtures/lookup_table.json",
			in: LookupTable{
				ID:          "7f042e5a-e4e8-47e7-8077-a0581411d4fd",
				Name:        "aws_account_metadata",
				Description: "Used to lookup business context on AWS accounts.",
				IDField:     "id",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/lookupTables/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/lookupTables/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateLookupTable(theT.in)
			assert.NoError(t, err)

			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/lookupTables/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateLookupTable(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      LookupTable
	}

	theTests := []convTest{
		{
			name:    "TestLookupTable",
			fixture: "fixtures/lookup_table.json",
			in: LookupTable{
				Name:            "aws_account_metadata",
				Description:     "Used to lookup business context on AWS accounts.",
				Active:          true,
				IDField:         "id",
				RowCount:        3,
				ForRuleEngine:   true,
				CreatedBy:       "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy:       "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				CreatedAt:       "2022-06-22T03:58:40.114Z",
				UpdatedAt:       "2022-06-22T03:58:40.349Z",
				DataLookupTable: DataLookupTable{},
				FetchRowsquery:  "SELECT id_field_value,data FROM upt_lookup_rows WHERE lookup_table_id = '7f042e5a-e4e8-47e7-8077-a0581411d4fd'",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Lookup Table information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/lookupTables/7f042e5a-e4e8-47e7-8077-a0581411d4fd"},
					LinkItem{Rel: "parent", Title: "Lookup Tables information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/lookupTables"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/lookupTables",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateLookupTable(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/lookupTables"], 1)
		})
	}
}
