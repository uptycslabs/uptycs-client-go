package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAlertRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   AlertRule
	}

	theTests := []convTest{
		{
			name: "TestAlertRule",
			in: AlertRule{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/alertRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteAlertRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/alertRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutAlertRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AlertRule
	}

	theTests := []convTest{
		{
			name:    "TestAlertRule",
			fixture: "fixtures/alertRuleBuilder.json",
			in: AlertRule{
				ID:          "9cde7195-ec0c-475e-a208-dbf81a32798a",
				Name:        "marcus test",
				Description: "marcus test",
				Code:        "test_marc",
				Type:        "sql",
				Rule:        "SELECT\n    'foo' as description,\n    'Low' as severity,\n    instance_id as asset_id,\n    now() as time,\n    'Alert' AS key,\n    'marcus test' as value\nFROM upt_cloud_instance_inventory_current\nWHERE key_name not like 'reddit2'\nLIMIT 1\n-- :to :from\n",
				Grouping:    "MITRE",
				Enabled:     true,
				GroupingL2:  "Impact",
				GroupingL3:  "T1560",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/alertRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateAlertRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/alertRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestGetAlertRule(t *testing.T) {

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
			name:    "TestAlertRuleBuilder",
			fixture: "fixtures/alertRuleBuilder.json",
			id:      "bcfe1cda-9eba-40fa-9686-9c1caea47732",
			out: AlertRule{
				ID:          "bcfe1cda-9eba-40fa-9686-9c1caea47732",
				SeedID:      "3a606e79-0e91-4e48-bcfe-dfb3798d0d33",
				Name:        "Login Profile Changes Detected_REDDIT",
				Description: "Changes detected in Login Profiles, which can result in Privilege Escalations",
				Code:        "AWS_THREAT_PRIV_ESC_4_REDDIT",
				Type:        "builder",
				Rule:        "Auto create alert rule for Login Profile Changes Detected_REDDIT",
				Grouping:    "ATTACK",
				Enabled:     true,
				Custom:      false,
				Throttled:   false,
				CreatedAt:   "2022-04-27T18:09:32.790Z",
				IsInternal:  false,
				AlertTags: []string{
					"ATTACK",
					"AWS",
					"Cloud",
					"IAM",
					"Privilege Escalation",
					"T1078",
					"THREAT",
					"elastalert",
				},
				CreatedBy: "e9b93444-a442-437f-82f6-6d65e9c787d3",
				UpdatedAt: "2022-04-27T18:09:32.790Z",
				//TimeSuppresionStart: null,
				//TimeSuppresionDuration: null,
				UpdatedBy:           "e9b93444-a442-437f-82f6-6d65e9c787d3",
				GroupingL2:          "Privilege Escalation",
				GroupingL3:          "T1078",
				Lock:                false,
				AlertNotifyInterval: 600,
				AlertNotifyCount:    1,
				Destinations: []AlertRuleDestination{
					AlertRuleDestination{
						ID:                 "caefdd0b-ca6a-4cec-b2ef-38cc45f73037",
						RuleID:             "bcfe1cda-9eba-40fa-9686-9c1caea47732",
						Severity:           "medium",
						DestinationID:      "f99515c0-9b6f-47df-a1a3-77bc1196b480",
						NotifyEveryAlert:   false,
						CloseAfterDelivery: false,
						CreatedAt:          "2022-04-28T20:07:57.800Z",
					},
				},
				AlertRuleExceptions: []RuleException{
					RuleException{
						ID:          "37e2f71c-5e5c-47b2-8d3b-6da4c198583d",
						RuleID:      "b4475f27-eaf4-44f3-9fa0-6dd0ca06089a",
						ExceptionID: "3c3dbd96-3f53-4123-8176-08b5c6d80db5",
						CreatedAt:   "2022-04-22T18:21:33.272Z",
						UpdatedAt:   "2022-04-22T18:21:33.277Z",
					},
					RuleException{
						ID:          "fa15d011-621b-4b18-a341-213a18659e03",
						RuleID:      "b4475f27-eaf4-44f3-9fa0-6dd0ca06089a",
						ExceptionID: "904eeeeb-cfb7-4e70-bfc1-955ec0b50971",
						CreatedAt:   "2022-04-22T18:21:33.277Z",
						UpdatedAt:   "2022-04-22T18:21:33.277Z",
					},
				},
				//AlertRuleQueries: [],
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Alert rule", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules/bcfe1cda-9eba-40fa-9686-9c1caea47732"},
					LinkItem{Rel: "parent", Title: "Alert rules", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/alertRules/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			alertRuleResp, err := c.GetAlertRule(AlertRule{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(alertRuleResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", alertRuleResp)
				t.Fail()
			}
		})
	}
}

func TestCreateAlertRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AlertRule
	}

	theTests := []convTest{
		{
			name:    "TestAlertRule",
			fixture: "fixtures/alertRuleCreate.json",
			in: AlertRule{
				Name:        "marcus test",
				Description: "marcus test",
				Code:        "test_marc",
				Type:        "sql",
				Rule:        "SELECT\n    'foo' as description,\n    'Low' as severity,\n    instance_id as asset_id,\n    now() as time,\n    'Alert' AS key,\n    'marcus test' as value\nFROM upt_cloud_instance_inventory_current\nWHERE key_name not like 'reddit2'\nLIMIT 1\n-- :to :from\n",
				Grouping:    "MITRE",
				Enabled:     true,
				GroupingL2:  "Impact",
				GroupingL3:  "T1560",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/alertRules",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateAlertRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/alertRules"], 1)
		})
	}
}
