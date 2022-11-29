package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetException(t *testing.T) {

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
			name:    "TestException",
			fixture: "fixtures/exception.json",
			id:      "ce67f12a-91d1-4a79-b0ee-a60501c5990b",
			out: Exception{
				ID:              "ce67f12a-91d1-4a79-b0ee-a60501c5990b",
				Name:            "Test Account Exception",
				Description:     "Enables Test Account Exclusion by Account ID",
				ExceptionType:   "alertRuleBuilder",
				CreatedAt:       "2022-05-12T21:32:44.834Z",
				CreatedBy:       "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt:       "2022-10-18T21:46:15.118Z",
				UpdatedBy:       "bccad1c2-7803-4341-bdda-27fe2c260bb2",
				TableName:       "aws_cloudtrail_events",
				IsGlobal:        true,
				Custom:          false,
				Disabled:        true,
				CloseOpenAlerts: false,
				Rule:            "{\"and\":[{\"caseInsensitive\":true,\"isDate\":false,\"isVersion\":false,\"isWordMatch\":false,\"name\":\"account_id\",\"not\":false,\"operator\":\"EQUALS\",\"value\":\"636776063332\"}]}",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Exception", Href: "/api/customers/11111111-1111-1111-1111-111111111111/exceptions/ce67f12a-91d1-4a79-b0ee-a60501c5990b"},
					LinkItem{Rel: "parent", Title: "Exceptions", Href: "/api/customers/11111111-1111-1111-1111-111111111111/exceptions"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/exceptions/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			exceptionResp, err := c.GetException(Exception{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(exceptionResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", exceptionResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteException(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   Exception
	}

	theTests := []convTest{
		{
			name: "TestException",
			in: Exception{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/exceptions/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteException(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/exceptions/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutException(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Exception
	}

	theTests := []convTest{
		{
			name:    "TestException",
			fixture: "fixtures/exception.json",
			in: Exception{
				ID:              "ce67f12a-91d1-4a79-b0ee-a60501c5990b",
				Name:            "Test Account Exception",
				Description:     "Enables Test Account Exclusion by Account ID",
				ExceptionType:   "alertRuleBuilder",
				CreatedAt:       "2022-05-12T21:32:44.834Z",
				CreatedBy:       "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt:       "2022-10-18T21:46:15.118Z",
				UpdatedBy:       "bccad1c2-7803-4341-bdda-27fe2c260bb2",
				TableName:       "aws_cloudtrail_events",
				IsGlobal:        true,
				Custom:          false,
				Disabled:        true,
				CloseOpenAlerts: false,
				Rule:            "{\"and\":[{\"caseInsensitive\":true,\"isDate\":false,\"isVersion\":false,\"isWordMatch\":false,\"name\":\"account_id\",\"not\":false,\"operator\":\"EQUALS\",\"value\":\"636776063332\"}]}",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Exception", Href: "/api/customers/11111111-1111-1111-1111-111111111111/exceptions/ce67f12a-91d1-4a79-b0ee-a60501c5990b"},
					LinkItem{Rel: "parent", Title: "Exceptions", Href: "/api/customers/11111111-1111-1111-1111-111111111111/exceptions"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/exceptions/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateException(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/exceptions/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateException(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Exception
	}

	theTests := []convTest{
		{
			name:    "TestException",
			fixture: "fixtures/exception.json",
			in: Exception{
				ID:              "ce67f12a-91d1-4a79-b0ee-a60501c5990b",
				Name:            "Test Account Exception",
				Description:     "Enables Test Account Exclusion by Account ID",
				ExceptionType:   "alertRuleBuilder",
				CreatedAt:       "2022-05-12T21:32:44.834Z",
				CreatedBy:       "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt:       "2022-10-18T21:46:15.118Z",
				UpdatedBy:       "bccad1c2-7803-4341-bdda-27fe2c260bb2",
				TableName:       "aws_cloudtrail_events",
				IsGlobal:        true,
				Custom:          false,
				Disabled:        true,
				CloseOpenAlerts: false,
				Rule:            "{\"and\":[{\"caseInsensitive\":true,\"isDate\":false,\"isVersion\":false,\"isWordMatch\":false,\"name\":\"account_id\",\"not\":false,\"operator\":\"EQUALS\",\"value\":\"636776063332\"}]}",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Exception", Href: "/api/customers/11111111-1111-1111-1111-111111111111/exceptions/ce67f12a-91d1-4a79-b0ee-a60501c5990b"},
					LinkItem{Rel: "parent", Title: "Exceptions", Href: "/api/customers/11111111-1111-1111-1111-111111111111/exceptions"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/exceptions",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateException(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/exceptions"], 1)
		})
	}
}
