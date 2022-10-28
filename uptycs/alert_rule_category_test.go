package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAlertRuleCategory(t *testing.T) {

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
			name:    "TestAlertRuleCategory",
			fixture: "fixtures/alertRuleCategories.json",
			id:      "b7c9c973-e2a3-4913-a755-919026267679",
			out: AlertRuleCategory{
				ID:        "79e61fc3-0557-431f-956f-80dee7a8a5fa",
				RuleID:    "7cf0b57d-1dea-4e91-b531-7661e5d97e7d",
				Name:      "foo",
				CreatedAt: "2022-10-25T19:00:54.275Z",
				CreatedBy: "f48f4c40-9c4a-47bb-9e3f-797d4deca92a",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Alert rule category information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRuleCategories/79e61fc3-0557-431f-956f-80dee7a8a5fa"},
					LinkItem{Rel: "parent", Title: "Alert rule category information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRuleCategories"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/alertRuleCategories/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			alertRuleCategoryResp, err := c.GetAlertRuleCategory(AlertRuleCategory{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(alertRuleCategoryResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", alertRuleCategoryResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteAlertRuleCategory(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   AlertRuleCategory
	}

	theTests := []convTest{
		{
			name: "TestAlertRuleCategory",
			in: AlertRuleCategory{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/alertRuleCategories/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteAlertRuleCategory(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/alertRuleCategories/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutAlertRuleCategory(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AlertRuleCategory
	}

	theTests := []convTest{
		{
			name:    "TestAlertRuleCategory",
			fixture: "fixtures/alertRuleCategories.json",
			in: AlertRuleCategory{
				Name: "different foo",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/alertRuleCategories/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateAlertRuleCategory(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/alertRuleCategories/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateAlertRuleCategory(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AlertRuleCategory
	}

	theTests := []convTest{
		{
			name:    "TestAlertRuleCategory",
			fixture: "fixtures/alertRuleCategories.json",
			in:      AlertRuleCategory{},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/alertRuleCategories",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateAlertRuleCategory(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/alertRuleCategories"], 1)
		})
	}
}
