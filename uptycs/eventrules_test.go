package uptycs

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetEventRules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/eventRules",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/eventRules.json")
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

	eventRulesResp, err := c.GetEventRules()
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
			in:   eventRulesResp,
			out: EventRules{
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventRules"},
					LinkItem{Rel: "parent", Href: "/api/customers/11111111-1111-1111-1111-111111111111"},
				},
				Items: []EventRule{
					EventRule{
						ID:            "ae79898a-7437-419d-88ff-e0e8d2c7902a",
						Name:          "badmonkey1652087888893",
						Description:   "test",
						Code:          "TESTCODE",
						Type:          "builder",
						Rule:          "builder",
						Grouping:      "cloud-compliance",
						Enabled:       true,
						Custom:        true,
						CreatedAt:     "2022-05-09T09:18:08.937Z",
						IsInternal:    false,
						EventTags:     []string{},
						CreatedBy:     "981e945c-010a-499d-b1ca-e68439e55fec",
						UpdatedAt:     "2022-05-09T09:18:08.937Z",
						UpdatedBy:     "00000000-0000-0000-0000-000000000000",
						GroupingL2:    "Impact",
						GroupingL3:    "Impact2",
						Score:         "9.0",
						Lock:          false,
						BuilderConfig: BuilderConfig{},
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Event generating rules", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventRules/ae79898a-7437-419d-88ff-e0e8d2c7902a"},
							LinkItem{Rel: "parent", Title: "Event generating ruless", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventRules"},
						},
					},
				},
				Offset: 0,
				Limit:  800,
			},
		},
	}

	for _, theT := range theTests {
		t.Run(theT.name, func(t *testing.T) {
			if !reflect.DeepEqual(theT.in, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Input: %v", theT.in)
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", eventRulesResp)
				t.Fail()
			}
		})
	}
}
