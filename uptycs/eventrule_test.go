package uptycs

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestPutEventRule(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   EventRule
	}

	theTests := []convTest{
		{
			name: "TestEventRule",
			in: EventRule{
				ID:          "b760d905-c161-43cd-8d44-d0ae8e1de1d5",
				Name:        "marc_is_awesomer",
				Description: "this is a test",
				Code:        "1651259159841CODE",
				Type:        "builder",
				Rule:        "builder",
				BuilderConfig: BuilderConfig{
					TableName:     "process_open_sockets",
					Added:         true,
					MatchesFilter: true,
					Filters: BuilderConfigFilter{
						And: []BuilderConfigFilter{
							{
								Name:     "remote_address",
								Operator: "MATCHES_REGEX",
								Value:    ArrayOrString{"^172.(1[6-9]|2[0-9]|3[01])|^10.|^192.168."},
								Not:      true,
							},
						},
					},
					Severity:   "low",
					Key:        "Test",
					ValueField: "pid",
				},
				EventTags: []string{
					"Tactic=Persistence",
					"Version=1.1",
					"Permissions Required=User",
				},
				Grouping:   "builderRules",
				GroupingL2: "Impact",
				GroupingL3: "T1560",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/eventRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.UpdateEventRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/eventRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestDeleteEventRule(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   EventRule
	}

	theTests := []convTest{
		{
			name: "TestEventRule",
			in: EventRule{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/eventRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteEventRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/eventRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateEventRule(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      EventRule
	}

	theTests := []convTest{
		{
			name:    "TestEventRule",
			fixture: "fixtures/eventRuleCreate.json",
			in: EventRule{
				Name:        "marc_is_awesomer",
				Description: "this is a test",
				Code:        "1651259159841CODE",
				Type:        "builder",
				Rule:        "builder",
				BuilderConfig: BuilderConfig{
					TableName:     "process_open_sockets",
					Added:         true,
					MatchesFilter: true,
					Filters: BuilderConfigFilter{
						And: []BuilderConfigFilter{
							{
								Name:     "remote_address",
								Operator: "MATCHES_REGEX",
								Value:    ArrayOrString{"^172.(1[6-9]|2[0-9]|3[01])|^10.|^192.168."},
								Not:      true,
							},
						},
					},
					Severity:   "low",
					Key:        "Test",
					ValueField: "pid",
				},
				EventTags: []string{
					"Tactic=Persistence",
					"Version=1.1",
					"Permissions Required=User",
				},
				Grouping:   "builderRules",
				GroupingL2: "Impact",
				GroupingL3: "T1560",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/eventRules",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateEventRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/eventRules"], 1)
		})
	}
}

func TestGetEventRule(t *testing.T) {

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
			name:    "TestBasicEventRule",
			fixture: "fixtures/eventRuleMinimal.json",
			id:      "e2f5af43-0044-40ef-b915-6dab1617166d",
			out: EventRule{
				ID:            "e2f5af43-0044-40ef-b915-6dab1617166d",
				CustomerID:    "111111111111-111111-11111-111111-111111111",
				SeedID:        "125b2c93-1dee-4a6c-ae80-45bcdf7000b0",
				Name:          "Process dropped script file on monitored locations - T1204.002 User Execution_LINUX",
				Description:   "Adversaries may drop scripts in system path to trick user into running them",
				Code:          "ATTACK_DROP_SCRIPT_T1204.002_LINUX_USER_EXECUTION_MALICIOUS_FILE",
				Type:          "builder",
				Rule:          "builder",
				Grouping:      "ATTACK",
				Enabled:       true,
				Custom:        false,
				CreatedAt:     "2021-06-02T14:43:26.899Z",
				IsInternal:    false,
				EventTags:     []string{"ATTACK", "Linux"},
				CreatedBy:     "00000000-0000-0000-0000-000000000000",
				UpdatedAt:     "2022-04-09T06:06:09.689Z",
				UpdatedBy:     "00000000-0000-0000-0000-000000000000",
				GroupingL2:    "Execution",
				GroupingL3:    "T1204",
				Score:         "0.0",
				Lock:          false,
				ScriptConfig:  ScriptConfig{},
				SQLConfig:     SQLConfig{},
				BuilderConfig: BuilderConfig{},
				//EventRuleExceptions: []string{},
				//Tags: []string{},
				//CloudResourceTags: []string{},
				//Exceptions: []string{},
				//Transformations: []string{},
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Event rule", Href: "/api/customers/111111111111-111111-11111-111111-111111111/eventRules/e2f5af43-0044-40ef-b915-6dab1617166d"},
					LinkItem{Rel: "parent", Title: "Event rules", Href: "/api/customers/111111111111-111111-11111-111111-111111111/eventRules"},
				},
			},
		},
		{
			name:    "TestFullEventRule",
			fixture: "fixtures/eventRuleFull.json",
			id:      "69bc42ba-d7c5-401d-b746-61afe5b372a2",
			out: EventRule{
				ID:          "69bc42ba-d7c5-401d-b746-61afe5b372a2",
				CustomerID:  "111111111111-111111-11111-111111-111111111",
				SeedID:      "125b2c93-1dee-4a6c-ae80-45bcdf7000b0",
				Name:        "Process dropped script file on monitored locations - T1204.002 User Execution_LINUX",
				Description: "Adversaries may drop scripts in system path to trick user into running them",
				Code:        "ATTACK_DROP_SCRIPT_T1204.002_LINUX_USER_EXECUTION_MALICIOUS_FILE",
				Type:        "builder",
				Rule:        "builder",
				Grouping:    "ATTACK",
				Enabled:     true,
				Custom:      false,
				CreatedAt:   "2021-06-02T14:43:26.899Z",
				IsInternal:  false,
				EventTags: []string{
					"ATTACK",
					"Linux",
					"Malicious File",
					"T1204.002",
					"User Execution",
					"Execution",
					"Endpoint",
					"process_file_events",
				},
				CreatedBy:  "00000000-0000-0000-0000-000000000000",
				UpdatedAt:  "2022-04-09T06:06:09.689Z",
				UpdatedBy:  "00000000-0000-0000-0000-000000000000",
				GroupingL2: "Execution",
				GroupingL3: "T1204",
				Score:      "0.0",
				Lock:       false,
				BuilderConfig: BuilderConfig{
					ID:            "69bc42ba-d7c5-401d-b746-61afe5b372a2",
					CustomerID:    "111111111111-111111-11111-111111-111111111",
					TableName:     "process_file_events",
					Added:         true,
					MatchesFilter: true,
					Filters: BuilderConfigFilter{
						And: []BuilderConfigFilter{
							BuilderConfigFilter{
								Or: []BuilderConfigFilter{
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/usr/local/sbin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/usr/local/bin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/usr/sbin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/usr/bin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/sbin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/bin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/usr/games/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/usr/local/games/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										Not:             false,
										Name:            "path",
										Value:           ArrayOrString{"/snap/bin/"},
										IsDate:          false,
										Operator:        "STARTS_WITH",
										CaseInsensitive: true,
									},
									BuilderConfigFilter{
										And: []BuilderConfigFilter{
											BuilderConfigFilter{
												Or: []BuilderConfigFilter{
													BuilderConfigFilter{
														Not:             false,
														Name:            "path",
														Value:           ArrayOrString{"/home/"},
														IsDate:          false,
														Operator:        "STARTS_WITH",
														IsVersion:       false,
														IsWordMatch:     false,
														CaseInsensitive: true,
													},
													BuilderConfigFilter{
														Not:             false,
														Name:            "path",
														Value:           ArrayOrString{"/root/"},
														IsDate:          false,
														Operator:        "STARTS_WITH",
														IsVersion:       false,
														IsWordMatch:     false,
														CaseInsensitive: true,
													},
												},
											},
											BuilderConfigFilter{
												Or: []BuilderConfigFilter{
													BuilderConfigFilter{
														Not:             false,
														Name:            "path",
														Value:           ArrayOrString{"/Downloads/"},
														IsDate:          false,
														Operator:        "CONTAINS",
														IsVersion:       false,
														IsWordMatch:     false,
														CaseInsensitive: true,
													},
													BuilderConfigFilter{
														Not:             false,
														Name:            "path",
														Value:           ArrayOrString{"/Download/"},
														IsDate:          false,
														Operator:        "CONTAINS",
														IsVersion:       false,
														IsWordMatch:     false,
														CaseInsensitive: true,
													},
													BuilderConfigFilter{
														Not:             false,
														Name:            "path",
														Value:           ArrayOrString{"/downloads/"},
														IsDate:          false,
														Operator:        "CONTAINS",
														IsVersion:       false,
														IsWordMatch:     false,
														CaseInsensitive: true,
													},
													BuilderConfigFilter{
														Not:             false,
														Name:            "path",
														Value:           ArrayOrString{"/download/"},
														IsDate:          false,
														Operator:        "CONTAINS",
														IsVersion:       false,
														IsWordMatch:     false,
														CaseInsensitive: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					Severity:        "medium",
					Key:             "Path",
					ValueField:      "path",
					AutoAlertConfig: AutoAlertConfig{},
				},
				//EventRuleExceptions: []string{},
				//Tags: []string{},
				//CloudResourceTags: []string{},
				//Exceptions: []string{},
				//Transformations: []string{},
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Event rule", Href: "/api/customers/111111111111-111111-11111-111111-111111111/eventRules/69bc42ba-d7c5-401d-b746-61afe5b372a2"},
					LinkItem{Rel: "parent", Title: "Event rules", Href: "/api/customers/111111111111-111111-11111-111111-111111111/eventRules"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/eventRules/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			eventRuleResp, err := c.GetEventRule(EventRule{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(eventRuleResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", eventRuleResp)
				t.Fail()
			}
		})
	}
}
