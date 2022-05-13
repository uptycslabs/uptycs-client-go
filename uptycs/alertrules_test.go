package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAlertRules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/alertRules",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/alertRules.json")
			if err != nil {
				t.Errorf(err.Error())
			}
			return fixture, err
		},
	)

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	alertRulesResp, err := c.GetAlertRules()
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
			in:   alertRulesResp,
			out: AlertRules{
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []AlertRule{
					AlertRule{
						ID:          "1d4720ce-19a9-4a03-bb3a-905717b8a60f",
						CustomerID:  "111111111111-111111-11111-111111-111111111",
						SeedID:      "30bd9a6f-259a-4506-b9da-80a8f392cced",
						Name:        "7zip.exe execution detected - T1560.001 - Archive via Utility - Windows",
						Description: "Adversaries may use 7zip utility to compress or encrypt exfilterated data in order to send it undetected to the attacker.",
						Code:        "ATTACK_ARCHIVE_VIA_UTILITY_DATA_T1560_001_WINDOWS_LOLBAS_7ZIP",
						Type:        "builder",
						Rule:        "Auto create alert rule for 7zip.exe execution detected - T1560.001 - Archive via Utility - Windows",
						Grouping:    "ATTACK",
						Enabled:     true,
						Custom:      false,
						Throttled:   false,
						CreatedAt:   "2021-06-02T14:44:59.818Z",
						IsInternal:  false,
						AlertTags: []string{
							"ATTACK",
							"Archive Collected Data",
							"Collection",
							"T1560",
							"T1560.001",
							"Windows",
							"Endpoint",
							"process_events",
						},
						CreatedBy: "00000000-0000-0000-0000-000000000000",
						UpdatedAt: "2022-04-09T06:01:41.060Z",
						//TimeSuppresionStart: null,
						//TimeSuppresionDuration: null,
						UpdatedBy:  "00000000-0000-0000-0000-000000000000",
						GroupingL2: "Collection",
						GroupingL3: "T1560",
						Lock:       false,
						//AlertNotifyInterval: null,
						//AlertNotifyCount: null,
						Destinations:        []Destination{},
						SQLConfig:           SQLConfig{},
						ScriptConfig:        ScriptConfig{},
						AlertRuleExceptions: []RuleException{},
						//AlertRuleQueries: [],
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Alert rule", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules/1d4720ce-19a9-4a03-bb3a-905717b8a60f"},
							LinkItem{Rel: "parent", Title: "Alert rules", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules"},
						},
					},
					AlertRule{
						ID:          "86d111f4-9fcb-4495-9cef-e2c35e63a47f",
						CustomerID:  "111111111111-111111-11111-111111-111111111",
						SeedID:      "AWS_THREAT_DEFENSE_EVASION_MODIFY_IAAS_1",
						Name:        "A VPC is deleted by an unknown user",
						Description: "A VPC is deleted by an unknown user",
						Code:        "AWS_THREAT_DEFENSE_EVASION_MODIFY_IAAS_1",
						Type:        "builder",
						Rule:        "Auto create alert rule for A VPC is deleted by an unknown user",
						Grouping:    "ATTACK",
						Enabled:     true,
						Custom:      false,
						Throttled:   false,
						CreatedAt:   "2022-01-14T21:16:36.900Z",
						IsInternal:  false,
						AlertTags: []string{
							"ATTACK",
							"AWS",
							"Defense Evasion",
							"IAAS",
							"MITRE",
							"MODIFY_IAAS",
							"Cloud",
							"T1578",
							"THREAT",
							"VPC",
						},
						CreatedBy: "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
						UpdatedAt: "2022-05-02T06:33:36.126Z",
						//TimeSuppresionStart: null,
						//TimeSuppresionDuration: null,
						UpdatedBy:  "00000000-0000-0000-0000-000000000000",
						GroupingL2: "Defense Evasion",
						GroupingL3: "T1578",
						Lock:       true,
						//AlertNotifyInterval: null,
						//AlertNotifyCount: null,
						Destinations:        []Destination{},
						SQLConfig:           SQLConfig{},
						ScriptConfig:        ScriptConfig{},
						AlertRuleExceptions: []RuleException{},
						//AlertRuleQueries: [],
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Alert rule", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules/86d111f4-9fcb-4495-9cef-e2c35e63a47f"},
							LinkItem{Rel: "parent", Title: "Alert rules", Href: "/api/customers/111111111111-111111-11111-111111-111111111/alertRules"},
						},
					},
				},
				Offset: 0,
				Limit:  999,
			},
		},
	}

	for _, theT := range theTests {
		t.Run(theT.name, func(t *testing.T) {
			if !reflect.DeepEqual(theT.in, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("  Actual: %v", alertRulesResp)
				t.Fail()
			}
		})
	}
}
