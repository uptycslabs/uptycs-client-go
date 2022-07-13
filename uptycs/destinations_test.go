package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetDestinations(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/destinations",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/destinations.json")
			if err != nil {
				t.Errorf(err.Error())
			}
			return fixture, err
		},
	)

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	destinationsResp, err := c.GetDestinations()
	if err != nil {
		t.Errorf(err.Error())
	}

	type convTest struct {
		name string
		in   interface{}
		out  interface{}
	}

	theTests := []convTest{
		{
			name: "thing",
			in:   destinationsResp,
			out: Destinations{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Notification destinations", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []Destination{
					Destination{
						ID:      "4c0dee1f-c19a-45fe-bf5d-fd031d6f694f",
						Name:    "tony-test",
						Type:    "email",
						Address: "tony.snook@reddit.com",
						//config: {},
						CreatedAt: "2021-07-12T21:13:09.778Z",
						CreatedBy: "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
						UpdatedAt: "2021-07-12T21:13:09.778Z",
						Enabled:   true,
						Default:   false,
						//Template: null,
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Notification destination", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations/4c0dee1f-c19a-45fe-bf5d-fd031d6f694f"},
							LinkItem{Rel: "parent", Title: "Notification destinations", Href: "/api/customers/111111111111-111111-11111-111111-111111111/destinations"},
						},
					},
				},
				Offset: 0,
				Limit:  1000,
			},
		},
	}

	for _, theT := range theTests {
		t.Run(theT.name, func(t *testing.T) {
			if !reflect.DeepEqual(theT.in, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("  Actual: %v", destinationsResp)
				t.Fail()
			}
		})
	}
}
