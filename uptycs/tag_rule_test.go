package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTagRule(t *testing.T) {

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
			name:    "TestTagRule",
			fixture: "fixtures/tagRule.json",
			id:      "56a72960-1673-418a-ac51-dbead6069742",
			out: TagRule{
				ID:           "56a72960-1673-418a-ac51-dbead6069742",
				Name:         "nginx",
				Description:  "nginx role",
				Query:        "select 'nginx' as tag from processes where name = 'nginx'",
				Source:       "realtime",
				RunOnce:      false,
				Interval:     3600,
				Platform:     "all",
				Enabled:      false,
				System:       true,
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				UpdatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2022-01-17T08:05:08.026Z",
				UpdatedAt:    "2022-04-09T05:59:46.347Z",
				ResourceType: "asset",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/tagRules/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			tagRuleResp, err := c.GetTagRule(TagRule{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(tagRuleResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", tagRuleResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteTagRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   TagRule
	}

	theTests := []convTest{
		{
			name: "TestTagRule",
			in: TagRule{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/tagRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteTagRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/tagRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutTagRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      TagRule
	}

	theTests := []convTest{
		{
			name:    "TestTagRule",
			fixture: "fixtures/tagRule.json",
			in: TagRule{
				ID:           "56a72960-1673-418a-ac51-dbead6069742",
				Name:         "nginx",
				Description:  "nginx role",
				Query:        "select 'nginx' as tag from processes where name = 'nginx'",
				Source:       "realtime",
				RunOnce:      true,
				Interval:     3600,
				Platform:     "all",
				Enabled:      false,
				System:       true,
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				UpdatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2022-01-17T08:05:08.026Z",
				UpdatedAt:    "2022-04-09T05:59:46.347Z",
				ResourceType: "asset",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/tagRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateTagRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/tagRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateTagRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      TagRule
	}

	theTests := []convTest{
		{
			name:    "TestTagRule",
			fixture: "fixtures/tagRule.json",
			in: TagRule{
				ID:           "56a72960-1673-418a-ac51-dbead6069742",
				Name:         "nginx",
				Description:  "nginx role",
				Query:        "select 'nginx' as tag from processes where name = 'nginx'",
				Source:       "realtime",
				RunOnce:      true,
				Interval:     3600,
				Platform:     "all",
				Enabled:      false,
				System:       true,
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				UpdatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2022-01-17T08:05:08.026Z",
				UpdatedAt:    "2022-04-09T05:59:46.347Z",
				ResourceType: "asset",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/tagRules",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateTagRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/tagRules"], 1)
		})
	}
}
