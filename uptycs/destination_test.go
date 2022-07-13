package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDestination(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
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
			name:    "TestDestination",
			fixture: "fixtures/destination.json",
			id:      "b7c9c973-e2a3-4913-a755-919026267679",
			out: Destination{
				ID:         "b7c9c973-e2a3-4913-a755-919026267679",
				CustomerID: "111111111111-111111-11111-111111-111111111",
				Name:       "#it-sec-alerts",
				Type:       "slack",
				Address:    "https://hooks.slack.com/services/wut/foo/asdf",
				//Config: {}, TODO
				CreatedAt: "2022-03-09T20:49:04.283Z",
				CreatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt: "2022-03-09T20:49:41.189Z",
				UpdatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				Enabled:   true,
				Default:   false,
				//Template: {}, TODO
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Notification destination", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations/b7c9c973-e2a3-4913-a755-919026267679"},
					LinkItem{Rel: "parent", Title: "Notification destinations", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/destinations/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			destinationResp, err := c.GetDestination(Destination{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(destinationResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", destinationResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteDestination(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   Destination
	}

	theTests := []convTest{
		{
			name: "TestDestination",
			in: Destination{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/destinations/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteDestination(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/destinations/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutDestination(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Destination
	}

	theTests := []convTest{
		{
			name:    "TestDestination",
			fixture: "fixtures/destinationCreate.json",
			in: Destination{
				ID:         "b7c9c973-e2a3-4913-a755-919026267679",
				CustomerID: "111111111111-111111-11111-111111-111111111",
				Name:       "#it-sec-alerts",
				Type:       "slack",
				Address:    "https://hooks.slack.com/services/wut/foo/asdf",
				//Config: {}, TODO
				CreatedAt: "2022-03-09T20:49:04.283Z",
				CreatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt: "2022-03-09T20:49:41.189Z",
				UpdatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				Enabled:   true,
				Default:   false,
				//Template: {}, TODO
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Notification destination", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations/b7c9c973-e2a3-4913-a755-919026267679"},
					LinkItem{Rel: "parent", Title: "Notification destinations", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/destinations/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateDestination(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/destinations/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateDestination(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Destination
	}

	theTests := []convTest{
		{
			name:    "TestDestination",
			fixture: "fixtures/destinationCreate.json",
			in: Destination{
				ID:         "b7c9c973-e2a3-4913-a755-919026267679",
				CustomerID: "111111111111-111111-11111-111111-111111111",
				Name:       "#it-sec-alerts",
				Type:       "slack",
				Address:    "https://hooks.slack.com/services/wut/foo/asdf",
				//Config: {}, TODO
				CreatedAt: "2022-03-09T20:49:04.283Z",
				CreatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt: "2022-03-09T20:49:41.189Z",
				UpdatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				Enabled:   true,
				Default:   false,
				//Template: {}, TODO
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Notification destination", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations/b7c9c973-e2a3-4913-a755-919026267679"},
					LinkItem{Rel: "parent", Title: "Notification destinations", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/destinations",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateDestination(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/destinations"], 1)
		})
	}
}
